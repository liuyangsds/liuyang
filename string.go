package liuyang

import (
	"math"
	"strconv"
	"strings"
	"unsafe"
)

// GetPrefix 获取字符串中指定字符的前缀(指定字符之前的字符串)，指定字符不存在时返回空字符串
// 样例1：
// str1 := "liuyang"
// tempStr1 := GetPrefix(str1, "_")
// 打印
// 空
// 样例2：
// str1 := "liuyang_sunlulu"
// tempStr1 := GetPrefix(str1, "_")
// 打印
// liuyang
func GetPrefix(s, sep string) string {
	pos := strings.Index(s, sep)
	if pos == NotFoundIndex {
		return ""
	}

	return s[:pos]
}

// GetPrefixV 获取字符串中指定字符的前缀(指定字符之前的字符串)，指定字符不存在时返回原字符串
// 样例1：
// str1 := "liuyang"
// tempStr1 := GetPrefixV(str1, "_")
// 打印
// liuyang
// 样例2：
// str1 := "liuyang_sunlulu"
// tempStr1 := GetPrefixV(str1, "_")
// 打印
// liuyang
func GetPrefixV(s, sep string) string {
	v := GetPrefix(s, sep)
	if v == "" {
		return s
	}

	return v
}

// GetSuffix 获取字符串中指定字符的后缀(指定字符之前的字符串)，指定字符不存在时返回空字符串
// 样例1：
// str1 := "liuyang"
// tempStr1 := GetSuffix(str1, "_")
// 打印
// 空
// 样例2：
// str1 := "liuyang_sunlulu"
// tempStr1 := GetSuffix(str1, "_")
// 打印
// liuyang
// //样例3：
// str3 := "_"
// tempStr2 := GetSuffix(str2, "_")
// 打印
// 空
func GetSuffix(s, sep string) string {
	pos := strings.Index(s, sep)
	if pos == NotFoundIndex {
		return ""
	}

	//获取字符串总长度
	length := len(s)
	//如果位置+1 大于等于 字符串总长度时，直接返回，这样判断可以避免1个字符时，取位置0再+1后，位置从1开始的问题
	index := pos + 1
	if index >= length {
		return ""
	}

	str := s[index:]

	return str
}

// GetSuffixV 获取字符串中指定字符的后缀(指定字符之前的字符串)，指定字符不存在时返回原字符串
// 样例1：
// str1 := "liuyang"
// tempStr1 := GetSuffixV(str1, "_")
// 打印
// liuyang
// 样例2：
// str1 := "liuyang_sunlulu"
// tempStr1 := GetSuffixV(str1, "_")
// 打印
// sunlulu
// 样例3：
// str2 := "_"
// tempStr2 := GetSuffixV(str2, "_")
// 打印
// _
func GetSuffixV(s, sep string) string {
	v := GetSuffix(s, sep)
	if v == "" {
		return s
	}

	return v
}

// 比较两个版本号的大小
// a大于b时，返回：1
// a小于b，返回：-1
// a和b相等时，返回：0
// 样例
// fmt.Println(CompareVersion("v2.11.9", "v2.10.8"))
// fmt.Println(CompareVersion("1.10.8", "1.19.7"))
// fmt.Println(CompareVersion("2.8.beta", "2.8"))
// 打印
// 1
// -1
// 0
func CompareVersion(a, b string) int {
	//获取两个字符串的后缀，如果没有指定字符v时，则返回原字符串
	a = GetSuffixV(a, "v")
	b = GetSuffixV(b, "v")

	if a == "" && b == "" {
		return 0
	}

	//将字符串切成切片
	var (
		arr1 = strings.Split(a, ".")
		arr2 = strings.Split(b, ".")
		diff int
	)

	//将两个切片的元素个数对齐，也就是元素少的补0
	//切片1的长度 - 切片2的长度
	diff = len(arr1) - len(arr2)
	for i := 0; i < diff; i++ {
		//如果切片2的长度比切片1的长度小时，则将切片2的元素补0
		arr2 = append(arr2, "0")
	}

	//切片2的长度 - 切片1的长度
	diff = len(arr2) - len(arr1)
	for i := 0; i < diff; i++ {
		//如果切片1的长度比切片2的长度小时，则将切片1的元素补0
		arr1 = append(arr1, "0")
	}

	//此时两个切片的元素个数相等，开始遍历中判断元素大小
	for i := 0; i < len(arr1); i++ {
		//将切片中的元素转为int类型
		ia, _ := strconv.Atoi(arr1[i])
		ib, _ := strconv.Atoi(arr2[i])

		if ia > ib {
			return 1
		}

		if ia < ib {
			return -1
		}
	}

	return 0
}

// 比较两个字符串的大小，如果是数字时，是不可靠的，如3会大于12.所以要特殊需求特殊对待。
// 如检测两个版本号的大小就使用特殊函数VersionCompare()进行检测。
// 字符串1大于字符串2时，返回：1
// 字符串1小于字符串2时，返回：-1
// 字符串1和字符串2相等时，返回：0
// 虽然和系统方法功能一样，但还是建议使用系统方法：strings.Compare()
// 样例
// n1 := test_new.CompareString("a", "b")
// n2 := test_new.CompareString("c", "b")
// n3 := test_new.CompareString("abcdea", "abcdeabc")
// n4 := test_new.CompareString("abcdea", "abcdea")
// 打印
// -1
// 1
// -1
// 0
func CompareString(str1, str2 string) int {
	lenS1 := len(str1)
	lenS2 := len(str2)
	for key, _ := range str1 {
		//当两个字符串中同位置字符相等时
		if str1[key] == str2[key] {
			lenS1--
			lenS2--
			//如果字符串1的长度为0时，并且字符串1的长度小于字符串2时
			if lenS1 == 0 && lenS1 < lenS2 {
				//fmt.Println("如果字符串1的长度为0时，并且字符串1的长度小于字符串2时")
				return -1
			}
			//如果字符串2的长度为0时，并且字符串2的长度小于字符串1时
			if lenS2 == 0 && lenS2 < lenS1 {
				//fmt.Println("如果字符串2的长度为0时，并且字符串2的长度小于字符串1时")
				return 1
			}
			continue //此处的continue是使用的最佳时机
		} else if str1[key] > str2[key] {
			//fmt.Println("这是字母1大于字母2")
			return 1
		} else {
			//fmt.Println("这是字母2大于字母1")
			return -1
		}
		//fmt.Println(str1[key],"==>",str2[key])
	}

	return 0
}

// 截取字符串，可以是所有字符，也包括汉字。
// 样例
// str := "abc刘阳def"
// aa := test_new.StringSubstr(str, 1, 5)
// fmt.Println(aa)
// 打印
// bc刘阳d
func StringSubstr(str string, startIndex int, length int) string {
	//这种方式是将汉字也转成1个字符
	rune_arr := []rune(str)

	//str := "abc刘阳def"
	//strlength := len(str)////(汉字算3个符字)
	runeLen := len(rune_arr) //(汉字算1个符字)
	//fmt.Println(strlength)//12
	//fmt.Println(runeLen)//8

	//如果长度大于总长度-起始位置
	if length > runeLen-startIndex {
		//长度=总长度-起始位置
		length = runeLen - startIndex
	}

	//必须+1才行
	length = length + 1

	return string(rune_arr[startIndex:length])
}

// 判断字符串是否为空，如果有空格的话，也视为空
// 样例1
// str := "abc刘阳def"
// aa := test_new.StringIsEmpty(str)
// fmt.Println(aa)
// 打印
// false
// 样例2
// str := "  \t "
// aa := test_new.StringIsEmpty(str)
// fmt.Println(aa)
// 打印
// true
func StringIsEmpty(str string) bool {

	//首先过滤掉字符串中的首尾空白字符
	tempStr := StringTrim(str)
	//再判断其长度是否大于0，如果大于0，说明不为空
	if len(tempStr) > 0 {
		return false
	}

	return true
}

// 过滤掉字符串中的首尾空白字符，包括空格、制表符、换页符等等。
func StringTrim(str string) string {
	//过滤掉字符串中的首尾空白字符，包括空格、制表符、换页符等等。
	tempStr := strings.Trim(str, " \f\n\r\t\v")
	return tempStr
}

// 得到一个英文字符串的ASCII码值的总和
func StringAscllSum(str string) int {
	sum := 0
	//fmt.Println(str)//liu yang
	for i := 0; i < len(str); i++ {
		//fmt.Print(str[i],"\t")//108	105	117	32	121	97	110	103
		sum += int(str[i])
	}
	//fmt.Println(sum)//793

	return sum
}

// string类型转int8
func StringToInt8(s string) int8 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt8 {
		return 0
	}

	return int8(n)
}

// string类型转uint8
func StringToUInt8(s string) uint8 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint8 {
		return 0
	}

	return uint8(n)
}

// string类型转int16
func StringToInt16(s string) int16 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt16 {
		return 0
	}

	return int16(n)
}

// string类型转uint16
func StringToUInt16(s string) uint16 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint16 {
		return 0
	}

	return uint16(n)
}

// string类型转int
func StringToInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt {
		return 0
	}

	return int(n)
}

// string类型转int32
func StringToInt32(s string) int32 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt32 {
		return 0
	}

	return int32(n)
}

// string类型转uint32
func StringToUInt32(s string) uint32 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint32 {
		return 0
	}

	return uint32(n)
}

// string类型转int64
func StringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt64 {
		return 0
	}

	return n
}

// string类型转uint64
func StringToUInt64(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}

	return n
}

// string类型转float32，虽然返回的是float64类型，但是以float32为精度截取的
func StringToFloat32(s string) float32 {
	//str := "123456.0123456789653"
	//fmt.Println(a1)//123456.015625
	ff, err := strconv.ParseFloat(s, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转换后的数值大于当前类型的最大值时，返回0
	if ff > math.MaxFloat32 {
		return 0
	}

	return float32(ff)
}

// string转float64
func StringToFloat64(s string) float64 {
	//str := "123456.0123456789653"
	//fmt.Println(a2)//123456.01234567896
	//值的注意：string转float时，只能转6位数以下的值，超过6位就会变成科学计数法，如下：
	//str := "123456"
	//fmt.Println(a2)//123456
	//str := "1234567"
	//fmt.Println(a2)//1.234567e+06
	ff, err := strconv.ParseFloat(s, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0 //这里不用返回该类型的最大值，返回0就代表转换失败
	}

	return ff
}

// 新增
// string转byte的高级写法
// 样例
// str := "abc刘阳def"
// aa := test_new.StringToBytes(str)
// fmt.Println(aa)
// bb := test_new.BytesToString(aa)
// fmt.Println(bb)
// 打印
// [97 98 99 229 136 152 233 152 179 100 101 102]
// abc刘阳def
func StringToBytes(s string) []byte {
	//return *(*[]byte)(unsafe.Pointer(&s))
	//上面的方式虽然正确，但是cap()却得不到正确的值，所以改用下面的方式
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// 新增
// byte转string的高级写法，bytes2string将字节片转换为字符串，无需内存分配
// 样例
// str := "abc刘阳def"
// aa := test_new.StringToBytes(str)
// fmt.Println(aa)
// bb := test_new.BytesToString(aa)
// fmt.Println(bb)
// 打印
// [97 98 99 229 136 152 233 152 179 100 101 102]
// abc刘阳def
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//===============================================================================

// 将带有百分号的字符串转成小数字符串。参数：带有百分号的数字或小数%(正负皆可)、保留小数位个数
// 样例1
// str := "23%"
// aa, ok := test_new.ScaleToString(str, 2)
// fmt.Println(ok)
// fmt.Println(aa)
// 打印
// true
// 0.23
// 样例2
// str := "2.3%"
// aa, ok := test_new.ScaleToString(str, 2)
// fmt.Println(ok)
// fmt.Println(aa)
// 打印
// true
// 0.02
// 样例3
// str := "2.3%"
// aa, ok := test_new.ScaleToString(str, 3)
// fmt.Println(ok)
// fmt.Println(aa)
// 打印
// true
// 0.023
func ScaleToString(str string, n int) (string, bool) {
	////如果后缀带有百分号“%”时
	//if strings.HasSuffix(str, "%") == true {
	//
	//	//获取除最后一个字符外的所有字符
	//	tempStr := str[0 : len(str)-1]
	//	//fmt.Println("值是：", tempStr) //15
	//
	//	//如果不全是数字时，直接返回
	//	//有符号(正负)检查小数，可输入0.000000，或输入小于1000000000(9个0)的正整数或小数(精确0至6位)，刘阳推荐
	//	if Check_s_decimal_6(tempStr) == false {
	//		//fmt.Println("如果不全是数字时，直接返回")
	//		return "", false
	//	}
	//
	//	//将string转成float64
	//	f_str := StringToFloat64(tempStr)
	//	f_str1 := f_str / 100
	//
	//	//将float64再转回str
	//	s_str := Float64ToString(f_str1, n)
	//
	//	return s_str, true
	//}

	//如果后缀带有百分号“%”时
	if strings.HasSuffix(str, "%") == true {
		//取%号前面的数值
		s := GetPrefixV(str, "%")
		//string 转 float64
		f := StringToFloat64(s)

		if f < 1 {
			return "", false
		}

		ff := f / 100

		//将float64再转回string
		ss := Float64ToString(ff, n)

		return ss, true

	}

	return "", false
}

// 将数字或小数(正负皆可)转成带有百分号的字符串。参数：数字或小数(正负皆可)、保留小数位个数。返回：15%
// 样例1
// str := "0.2356"
// aa := test_new.StringToScale(str, 2)
// fmt.Println(aa)
// 打印
// 23.56%
// 样例2
// str := "0.23"
// aa := test_new.StringToScale(str, 2)
// fmt.Println(aa)
// 打印
// 23.00%
// 样例3
// str := "0.23"
// aa := test_new.StringToScale(str, 0)
// fmt.Println(aa)
// 打印
// 23%
func StringToScale(str string, n int) string {
	f := StringToFloat64(str)
	ff := f * 100

	s_f := Float64ToString(ff, n)
	s_f += "%" //连接百分号

	return s_f
}

//------------------------------------------------旧函数
