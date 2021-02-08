package liuyang

import "fmt"

func Number_test()  {
	for i := 0; i > -10; i-- {
		aa := Number_m(i, 4)
		fmt.Println("i的值是：",i,"返回值是：",aa)
	}

}

//数字取模，参数1为自然数，参数2为模
func Number_m(n int, m int) int {

	//如果参数大于等于0时
	if n >= 0 {
		return n % m
	}

	return Number_m(n + m, m)//这里必须要写返回才行，不然全得到的全是0
}
