package liuyang

import (
	"fmt"
	"math"
	"strconv"
)

//方式1，将float64类型的小数四舍六入五成双后，得到一个字符串，且保留指定位数的小数
func Float64ToString(ff float64, n int) string {
	//参数n表示：
	//0，不保留小数，校例：
	//var tempF float64 = 12345678.51
	//var tempF float64 = 12345678.50 //第2位是0则第1位的5就不向前进1了
	//保留0位后： 12345679
	//保留0位后： 12345678
	//1，保留小数点后1位
	//2，保留小数点后2位
	//3，保留小数点后3位
	//-1，保留小数点后所有数字
	//参数bitSize:64表示以float64为精度进行保留小数，可保留整数加小数共计15位，比32位要好很多，推荐使用
	if n < 0 { //如果小于0时，也就是-1...-n的时候，将其赋值为-1。意思是保留所有小数位。
		n = -1
	}
	//参数说明：
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	//这里的'f'为byte类型，所以必须得是单引号，而双引号则是string类型，会报错的。
	str := strconv.FormatFloat(ff, 'f', n, 64)
	return str
}

//方式2，将float64类型的小数四舍六入五成双后，得到一个固定保留2位小数的字符串
func Float64ToString2(ff float64) string {
	str := fmt.Sprintf("%.2f", ff)
	return str
}

//将float64转成指定小数位的float64
func Float64ToFloat64(ff float64, n int) float64 {
	//1，先将float类型转成string类型，并指定保留的小数个数
	str := Float64ToString(ff, n)
	//2，再将string类型转成float64
	ss := StringToFloat64(str)

	return ss
}

//float64转int64
func Float64ToInt64(ff float64) int64 {
	f := math.Trunc(ff)

	return int64(f)
}
