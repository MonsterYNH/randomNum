package testExam

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_RandomNumIsPositiveNumber(t *testing.T) {
	Convey("测试生成的10000个随机数为正数", t, func() {
		var isNegativeNumber bool
		for i := 0; i < 10000; i++ {
			if num:= PsudoEncrypt(); num < 0 {
				isNegativeNumber = true
				break
			}
		}
		So(isNegativeNumber, ShouldBeFalse)
	})
}

func Test_RondomNumIsRandom(t *testing.T) {
	Convey("测试生成10000个随机数不重复", t, func() {
		numMap := make(map[int32]int)
		var isDuplicate bool
		for i := 0; i < 10000; i++ {
			// 为了保证随机数种子每次都不相同
			time.Sleep(time.Microsecond)
			num := PsudoEncrypt()
			if _, exist := numMap[num]; exist {
				isDuplicate = true
				break
			} else {
				numMap[num] = 1
			}
		}
		So(isDuplicate, ShouldBeFalse)
	})
}
