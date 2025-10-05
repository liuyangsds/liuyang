package liuyang

import (
	"math"
)

// 定点数-编译(将普通数值转为定点数)--->参数1：数值、参数2：保留的小数位数
// 样例1
// f := 2.32
// n := test_new.FixedPointEncode(f, 2)
// fmt.Println(n)
// 打印
// 232
// 样例2
// f := 19.988
// n := test_new.FixedPointEncode(f, 2)
// fmt.Println(n)
// 打印
// 1999
// 样例3
// f := 2.36
// n := test_new.FixedPointEncode(f, 2)
// fmt.Println(n)
// 打印
// 236
func FixedPointEncode(f float64, digit int) int {
	//以保留的小数位数乘以10，得到10的倍数：10、100、1000、10000。。。
	offset := math.Pow10(digit)

	//实际数值 * 倍数
	f = f * offset
	//fmt.Println(f) //231.99999999999997

	//直接转int会丢失小数后面的值，也就是没有四舍五入
	//如：2.32在传值时会变成231.99999999999997，乘100后得231.99999999999997，丢掉小数后得231
	//n := int(f)

	//float转int函数是自定义封装函数，里面有一层是float转string，再由string转int，这样会四会五入。
	n := Float64ToInt(f)

	return n
}

// 定点数-解析(将定点数还原为普通数值)--->参数1：数值、参数2：保留的小数位数
// 样例1
// f := -2.32
// n := test_new.FixedPointEncode(f, 2)
// fmt.Println("取两位小数后的定点数：", n)
// s := test_new.FixedPointDecode(n, 2)
// fmt.Println("恢复两位小数后的定点数：", s)
// 打印
// 取两位小数后的定点数： -232
// 恢复两位小数后的定点数： -2.32
// 样例2
// f := 19.988
// n := test_new.FixedPointEncode(f, 2)
// fmt.Println("取两位小数后的定点数：", n)
// s := test_new.FixedPointDecode(n, 2)
// fmt.Println("恢复两位小数后的定点数：", s)
// 打印
// 取两位小数后的定点数： 1999
// 恢复两位小数后的定点数： 19.99
func FixedPointDecode(n int, digit int) string {

	//将int型数值转为float64
	f := float64(n)

	//以保留的小数位数乘以10，得到10的倍数：10、100、1000、10000。。。
	offset := math.Pow10(digit)

	//实际数值 / 倍数
	ff := f / offset

	//float 四舍五入后转 string
	s := Float64ToString(ff, digit)

	return s
}
