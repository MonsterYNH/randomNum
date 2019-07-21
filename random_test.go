package testExam

import (
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_RandomNumIsPositiveNumber(t *testing.T) {
	Convey("测试生成的10000个随机数为正数", t, func() {
		var isNegativeNumber bool
		for i := 0; i < 10000; i++ {
			if num, _ := PsudoEncrypt(0x800000000000); num < 0 {
				isNegativeNumber = true
				break
			}
		}
		So(isNegativeNumber, ShouldBeFalse)
	})
}

func Test_RondomNumIsRandom(t *testing.T) {
	Convey("测试生成10000个随机数不重复", t, func() {
		numMap := make(map[int64]int)
		var isDuplicate bool
		for i := 0; i < 10000; i++ {
			// 为了保证随机数种子每次都不相同
			time.Sleep(time.Microsecond)
			num, _ := PsudoEncrypt(0x800000000000)
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

func Test_RondomNumIsFixedLength(t *testing.T) {
	Convey("测试生成10000个随机数的长度固定", t, func() {
		var isFixedLength bool
		for i := 0; i < 10000; i++ {
			num, _ := PsudoEncrypt(0x800000000000)
			if len(strconv.FormatInt(num, 10)) != 15 {
				isFixedLength = true
				break
			}

		}
		So(isFixedLength, ShouldBeFalse)
	})
}

//func TestGengeratUserNameWithInfo(t *testing.T) {
//	id, _ := GengeratUserNameWithInfo(AD_HIGH, INFO_VIP, LOC_BEIJING, 0x800000000000)
//	fmt.Printf("%+v", id)
//}
