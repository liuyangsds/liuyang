package liuyang

//定点数相关函数

//获取定点数-整型转浮点型后除100
func GetFixedPointIntChu100(n int64) float64 {
	//fmt.Println("参数值：", n)
	f := float64(n) / 100
	//fmt.Println("转换后的定点数：", f)

	return f
}

//获取定点数-浮点型乘100后转整型
func GetFixedPointFloatCheng100(f float64) int64 {
	//fmt.Println("参数值：", f)
	//截取小数点后两位
	f_f := Float64ToFloat64(f, 2)
	//fmt.Println("截取两位后：", f_f)
	f_f = f_f * 100
	i_f := Float64ToInt64(f_f)
	//fmt.Println("转换后的定点数：", i_f)

	return i_f
}

//获取定点数-整型乘100后的整型
func GetFixedPointIntCheng100(n int64) int64 {
	n = n * 100

	return n
}
