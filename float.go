package liuyang

import (
	"math"
	"strconv"
)

//将float64类型的小数四舍六入五成双后，得到一个字符串
func Float64ToString(f float64,n int) string  {
	//参数n表示：
	//0，取整数，不要后面的小数点部分，但是四舍五入后2.50变成了3，1.50变成了2,这是不理想的，所以0时要单独判断
	//1，保留小数点后1位
	//2，保留小数点后2位
	//3，保留小数点后3位
	//-1，保留小数点后所有数字
	//参数bitSize:64表示以float64为精度进行保留小数，可保留整数加小数共计15位，比32位要好很多，推荐使用
	if n == 0 {
		//首先取整
		newf := math.Trunc(f)
		//再强转int，为什么不直接强转呢？因为那样太强硬了，对程序不友好，这样会和缓的将float64类型的小数先取整，再将整数部分转int比较好
		num := int64(newf)//将float64类型的整数强转int64
		str := strconv.FormatInt(num,10)//再将int64转成string类型，10代表进制。
		return str
	} else if n < 0 {//如果小于0时，也就是-1...-n的时候，将其赋值为-1。意思是保留所有小数位。
		n = -1
	}
	//参数说明：
	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	//prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	str := strconv.FormatFloat(f,'f',n,64)
	return str
}
