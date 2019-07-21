package testExam

import (
	"errors"
	"math"
	"time"
)

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
