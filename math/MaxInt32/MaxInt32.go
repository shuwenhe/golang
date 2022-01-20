package MaxInt32

import (
	"fmt"
	"math"
)

const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func min_int32() {
	fmt.Println("min_int32 = ", MinInt32)
}

func max_int32() {
	fmt.Println("math.MaxInt32 = ", math.MaxInt32)
}
