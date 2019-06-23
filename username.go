package randgo

import (
	"fmt"
	"math/rand"
	"sync"
)

//思路：将形容词、人名、和地名放在三个不同的数组中，然后以三个数组中的最大长度-1中选中3个数生成全排列，储存在数组中，以这个数组的最大长度-1使用
//sort.Int63n()生成随机数，选出排列数组中的排列，然后看对应的数字是否超出形容词、人名、地名数组的最大长度，超出则将其从数组中删除，以新的长度重
//新随机。不超出则将其组合成用户名，并将对应的排列从数组删除。或者使用map储存已经使用或者无效的排列的下标号。实际生产中如果词语比较多，可以将形容
//词、人名、地名以 序号/形容词/人名/地名之类的形式储存在数据库中，节约内存的使用，使代码更加干净。
//后续优化思路：以更加节约内存的方式储存生成的排列（可能会牺牲一部分生成用户名的速度），此外目前的方法一旦生成排列并开始运行就无法轻易的扩展形容词、
//人名、地名的数组或数据库表，需要一个扩展性更好的实现思路

//今天上午还有其他事情，没有时间写出实现，下午或者晚上有时间可能会提交具体的实现。
//PS:那个算法虽然知道是使用的异或、移位操作的不可逆原理，但还是写不大来。。。

//排列组合算法参考https://blog.csdn.net/books1958/article/details/46861341

type array struct {
	array [][]int
	sync.Mutex
	maxlen *maxlen
}

type maxlen struct {
	n int
	m int
	l int
}

func (a *array) GenerateUserName() []int {
	a.Lock()
	defer a.Unlock()

	var result []int
	r := rand.Int63n(int64(len(a.array)))
	for len(a.array) > 1 {
		if a.array[r][0] < a.maxlen.n && a.array[r][1] < a.maxlen.m && a.array[r][2] < a.maxlen.l {
			result = a.array[r]
			a.deleteFromArray(int(r))
			break
		}
		a.deleteFromArray(int(r))
		r = rand.Int63n(int64(len(a.array)))
	}

	return result
}

func (a *array) deleteFromArray(n int) {
	switch {
	case n == 0:
		a.array = a.array[1:]
	case n == len(a.array)-1:
		a.array = a.array[:n]
	default:
		a.array = append(a.array[:n], a.array[n+1:]...)
	}
}

func GenerateArray(n, m, l int) *array {
	max := n
	if m > n && m > l {
		max = m
	} else if l > n {
		max = l
	}

	maxlen := &maxlen{
		n: n,
		m: m,
		l: l,
	}

	nums := make([]int, 0)
	for i := 0; i < max; i++ {
		nums = append(nums, i)
	}
	// log.Println(nums)
	// log.Println(pailieResult(max, 3))

	a := new(array)
	a.array = pailieResult(nums, max, 3)
	a.maxlen = maxlen

	// log.Println(a.array)
	return a
}

func zuheResult(n int, m int) [][]int {
	if m < 1 || m > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}

	result := make([][]int, 0, mathZuhe(n, m))
	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}

	result = addTo(result, indexs)
	for {
		find := false
		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true

				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)

				break
			}
		}

		if !find {
			break
		}
	}

	return result
}

func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)

	return arr
}

func moveOneToLeft(leftNums []int) {
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}

	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

func findNumsByIndexs(nums []int, indexs [][]int) [][]int {
	if len(indexs) == 0 {
		return [][]int{}
	}

	result := make([][]int, len(indexs))

	for i, v := range indexs {
		line := make([]int, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}

	return result
}

func pailieResult(nums []int, n, m int) [][]int {
	zuhe := findNumsByIndexs(nums, zuheResult(n, m))

	result := make([][]int, 0)
	for _, v := range zuhe {
		p := quanPailie(v)
		result = append(result, p...)
	}

	return result
}

func quanPailie(nums []int) [][]int {
	COUNT := len(nums)
	if COUNT == 0 || COUNT > 10 {
		panic("Illegal argument. nums size must between 1 and 9.")
	}

	if COUNT == 1 {
		return [][]int{nums}
	}

	return insertItem(quanPailie(nums[:COUNT-1]), nums[COUNT-1])
}

func insertItem(res [][]int, insertNum int) [][]int {
	result := make([][]int, len(res)*(len(res[0])+1))

	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			result[index] = insertToSlice(v, i, insertNum)
			index++
		}

		result[index] = append(v, insertNum)
		index++
	}

	return result
}

//将元素value插入到数组nums中索引为index的位置
func insertToSlice(nums []int, index int, value int) []int {
	result := make([]int, len(nums)+1)
	copy(result[:index], nums[:index])
	result[index] = value
	copy(result[index+1:], nums[index:])

	return result
}

func mathZuhe(n int, m int) int {
	return jieCheng(n) / (jieCheng(n-m) * jieCheng(m))
}

//阶乘
func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func mathPailie(n int, m int) int {
	return jieCheng(n) / jieCheng(n-m)
}
