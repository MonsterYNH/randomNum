package testExam

import (
	"errors"
	"fmt"
	"math"
	"time"
)


const (
	// 形容词
	AD_HIGH = iota
	// ...
	// 身份信息
	INFO_VIP
	// ...
	// 地点
	LOC_BEIJING
	// ...
)

var (
	adConvertMap map[int]string
	adMap map[int]int8
	infoConvertMap map[int]string
	infoMap map[int]int8
	locConvertMap map[int]string
	locMap map[int]int8
)

func init() {
	adConvertMap = make(map[int]string)
	adMap = make(map[int]int8)
	adConvertMap[AD_HIGH] = "fancy"
	adMap[AD_HIGH] = 0x00

	infoConvertMap = make(map[int]string)
	infoMap = make(map[int]int8)
	infoConvertMap[INFO_VIP] = "blacksmith"
	infoMap[INFO_VIP] = 0x00

	locConvertMap = make(map[int]string)
	locMap = make(map[int]int8)
	locConvertMap[LOC_BEIJING] = "beijing"
	locMap[LOC_BEIJING] = 0x00
}
func PsudoEncrypt(base int64) (int64, error) {
	value := int32(time.Now().UnixNano())
	value = (value ^ value>>31) - (value>> 31)
	l1 := value >> 16 & 0xffff
	r1 := value & 0xffff

	var l2, r2 int32
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int32(math.Floor((float64((1366 * r1 + 150889) % 714025) / 714025.0) * 32767))
		l1 = l2
		r1 = r2
	}

	num := int64((r1<<16)+l1)
	if (base >> 2) < num {
		return -1, errors.New("base num is too small to create a random num")
	} else {
		return base+num, nil
	}
}

type ID struct {
	Code int64
	Str string
}

func GengeratUserNameWithInfo(ad, info, loc int, base int64) (*ID, error) {
	// 前24位中，前8位代表形容词,中8位代表身份信息，后8位代表地方
	var id int64
	adCode := int64(adMap[ad]) << (64-8)
	infoCode := int64(infoMap[info]) << (64-16)
	locCode := int64(locMap[loc]) << (64-24)
	id = adCode + infoCode + locCode
	// 后40位为随机数区域保证唯一
	num, err := PsudoEncrypt(base)
	if err != nil {
		return nil, nil
	}
	id += num
	return &ID{
		Code: id,
		Str: fmt.Sprintf("%s %s from %s", adConvertMap[ad], infoConvertMap[info], locConvertMap[loc]),
	}, nil
}
