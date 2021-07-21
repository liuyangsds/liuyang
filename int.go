package liuyang

import "strconv"

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func UInt64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}