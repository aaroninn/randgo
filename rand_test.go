package randgo

import (
	"log"
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
	})

}
