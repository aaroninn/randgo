package randgo

import (
	"log"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_encrypt(t *testing.T) {
	Convey("test Rand", t, func() {
		generated := make(map[int]bool)

		for i := 0; i < 100000; i++ {
			e := Rand()
			_, ok := generated[e]
			log.Println(e)
			So(e, ShouldBeGreaterThan, 0)
			So(ok, ShouldBeFalse)
			generated[e] = true
		}

		Convey("test RandN", func() {
			generated := make(map[int]bool)

			for i := 0; i < 1000000; i++ {
				e := RandN(18)
				_, ok := generated[e]
				So(e, ShouldBeGreaterThan, int(math.Pow10(17)))
				So(e, ShouldBeLessThan, int(math.Pow10(18))-1)
				So(ok, ShouldBeFalse)
				generated[e] = true
			}
		})
	})

}
