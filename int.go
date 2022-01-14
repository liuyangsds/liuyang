package liuyang

import "strconv"

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func UInt64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}

//============================================
func Int32ToString(n int32) string {
	x := int64(n)
	return strconv.FormatInt(x, 10)
}

func UInt32ToString(n uint32) string {
	x := uint64(n)
	return strconv.FormatUint(x, 10)
}

//============================================
func Int16ToString(n int16) string {
	x := int64(n)
	return strconv.FormatInt(x, 10)
}

func UInt16ToString(n uint16) string {
	x := uint64(n)
	return strconv.FormatUint(x, 10)
}

//============================================
func Int8ToString(n int8) string {
	x := int64(n)
	return strconv.FormatInt(x, 10)
}

func UInt8ToString(n uint8) string {
	x := uint64(n)
	return strconv.FormatUint(x, 10)
}

//============================================
