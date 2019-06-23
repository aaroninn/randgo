package randgo

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GenerateUserName(t *testing.T) {
	Convey("test generateusername", t, func() {
		array := GenerateArray(4, 3, 2)
		//排列数 4！/（4-3）！乘以组合数4！/（（4-3)!*3!)
		So(len(array.array), ShouldEqual, 24)
		for i := 0; i < 6; i++ {
			result := array.GenerateUserName()
			log.Println("result is ", result)
			So(len(result), ShouldEqual, 3)
			So(result[0], ShouldBeLessThan, 4)
			So(result[1], ShouldBeLessThan, 3)
			So(result[2], ShouldBeLessThan, 2)
		}
	})
}
