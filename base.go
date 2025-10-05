package liuyang

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ====================================10进制转成2、8、16进制========================================

// 进制转换，10转2
func Base10To2[T byte | int | int16 | int32 | int64 | uint16 | uint32 | uint64](n T) string {
	return strconv.FormatInt(int64(n), 2)
}

// 进制转换，10转8
func Base10To8[T byte | int | int16 | int32 | int64 | uint16 | uint32 | uint64](n T) string {
	return strconv.FormatInt(int64(n), 8)
}

// 进制转换，10转16
func Base10To16[T byte | int | int16 | int32 | int64 | uint16 | uint32 | uint64](n T) string {
	return strconv.FormatInt(int64(n), 16)
}

// ====================================10进制转成2、8、16进制========================================

// ====================================2、8、16进制转成10进制========================================

// 进制转换，2816转10--->可根据0b、0o、0x前缀自动判断想要的进制计算
// ParseInt函数说明：base为0时，代表参数可以传以0b、0o、0x开头的字符串，如果base大于0就不允许传特殊字符。
// ab2 := test.Base0To10("0b10110") //22
// ab8 := test.Base0To10("0o26")    //22
// ab16 := test.Base0To10("0x16")   //22
func Base2816To10(s string) int64 {
	n, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0
	}

	return n
}

// 进制转换，2转10--->注意：不允许带前缀0b，因为ParseInt函数中已指定了2进制，如想带前缀，请使用Base0To10()函数
// 样例1
// str := "0b10011"
// bb := test_new.Base2To10(str)
// 打印
// 报错：strconv.ParseInt: parsing "0b10011": invalid syntax
// 样例2
// str := "10011"
// bb := test_new.Base2To10(str)
// 打印
// 19
func Base2To10(s string) int64 {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}

	return n
}

// 进制转换，8转10--->注意：不允许带前缀0o，因为ParseInt函数中已指定了8进制，如想带前缀，请使用Base0To10()函数
// 样例1
// str := "0o177"
// bb := test_new.Base8To10(str)
// 打印
// 报错：strconv.ParseInt: parsing "0o177": invalid syntax
// 样例2
// str := "177"
// bb := test_new.Base8To10(str)
// 打印
// 127
func Base8To10(s string) int64 {
	n, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		return 0
	}

	return n
}

// 进制转换，16转10--->注意：不允许带前缀0x，因为ParseInt函数中已指定了16进制，如想带前缀，请使用Base0To10()函数
// 样例1
// str := "0x177"
// bb := test_new.Base16To10(str)
// 打印
// 报错：strconv.ParseInt: parsing "0x177": invalid syntax
// 样例2
// str := "177"
// bb := test_new.Base16To10(str)
// 打印
// 375
func Base16To10(s string) int64 {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0
	}

	return n
}

// ====================================2、8、16进制转成10进制========================================

// 打印数组，将其元素10进制转成2、8、16进制并以指定字符连接
// 参数1：要打印的数组、参数2：指定要转成的进制、参数3：转换时的前缀模式、参数4：转换后每个元素的连接符
// 校例：([]byte, 16, 0, ",")，0代表前缀是标准的2(ob)、8(0o）、16(0x)，并将1-9自动补0变为01-09。
// 2进制-样例1
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 2, 0)
// 打印
// 11,1001,11001,1100,100100,1010
// 2进制-样例2
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 2, 1)
// 打印
// 0b11,0b1001,0b11001,0b1100,0b100100,0b1010
// 8进制-样例1
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 8, 0)
// 打印
// 3,11,31,14,44,12
// 8进制-样例2
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 8, 1)
// 打印
// 0o3,0o11,0o31,0o14,0o44,0o12
// 16进制-样例1
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 16, 0)
// 打印
// 03,09,19,0c,24,0a
// 16进制-样例2
// arr := []int{3, 9, 25, 12, 36, 10}
// str := test_new.Print10To2816(arr, ",", 16, 1)
// 打印
// 0x03,0x09,0x19,0x0c,0x24,0x0a
func Print10To2816[T byte | int | int16 | int32 | int64 | uint16 | uint32 | uint64](arr []T, sep string, base int, flag uint8) string {
	var str = ""      //总字符串，默认空
	var preX = ""     //补字符0b、0o、0x,默认空
	var base_str = "" //进制不同，连接的字符也不同
	switch base {
	case 2:
		base_str = "0b"
	case 8:
		base_str = "0o"
	case 16:
		base_str = "0x"
	default:
		base_str = ""
	}
	if flag == 0 {
		preX = ""
	} else if flag == 1 {
		preX = base_str
	}
	for key, value := range arr {
		//将10进制数转成16进制字符串
		cardData := strconv.FormatInt(int64(value), base)

		//如果要转为16进制时，需考虑补0的情况
		if base == 16 {
			//如果长度为1，也就是0到9时，补0，这样方便看，不然会是0x1,0x2，远没有0x01,0x02好看。
			if len(cardData) == 1 {
				cardData = "0" + cardData
			}
		}

		if key < len(arr)-1 {
			str += preX + cardData + sep
		} else {
			str += preX + cardData
		}
	}

	return str
}

// ====================================传统计算方式封装的函数========================================

// 将10进制转2进制后得到字符串(传统计算方式)
func Base10To2Custom(n int64) string {
	if n < 0 {
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s) //将000111这样的值一个一个字符串拼接给临时变量s
	}
	return s
}

// 将10进制转8进制后得到字符串(传统计算方式)
func Base10To8Custom(n int64) string {
	if n < 0 {
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 8 {
		m := q % 8
		s = fmt.Sprintf("%v%v", m, s) //将000111这样的值一个一个字符串拼接给临时变量s
	}
	return s
}

// 将10进制数值转成16进制字符串(传统计算方式)
// 样例1
// aa := test_new.Base10To16Custom(15)
// fmt.Println(aa)
// bb := test_new.Base16To10Custom(aa)
// fmt.Println(bb)
// 打印
// F
// 15
// 样例2
// aa := test_new.Base10To16Custom(1024)
// fmt.Println(aa)
// bb := test_new.Base16To10Custom(aa)
// fmt.Println(bb)
// 打印
// 400
// 1024
func Base10To16Custom(n int64) string {
	if n < 1 {
		return "0"
	}

	//16进制
	var b int64 = 16

	s := ""
	for i := n; i > 0; i = i / b {
		m := i % b
		if m > 9 && m < b { //如果余数是10到15之间的数时
			temp := ""
			switch m {
			case 10:
				temp = "A"
			case 11:
				temp = "B"
			case 12:
				temp = "C"
			case 13:
				temp = "D"
			case 14:
				temp = "E"
			case 15:
				temp = "F"
			}

			//s = fmt.Sprintf("%v%v", temp, s) //将000111这样的值一个一个字符串拼接给临时变量s
			s = temp + s //将000111这样的值一个一个字符串拼接给临时变量s
			//fmt.Println("循环内--->i值：", i, "，m值：", m, "，s值：", s)
			continue //跳出，下面的代码不执行，很关键，一定要写才行。
		}

		//s = fmt.Sprintf("%v%v", m, s) //将000111这样的值一个一个字符串拼接给临时变量s
		//s = strconv.Itoa(m) + s //将000111这样的值一个一个字符串拼接给临时变量s
		s = Int64ToString(m) + s //将000111这样的值一个一个字符串拼接给临时变量s
		//fmt.Println("循环内--->i值：", i, "，m值：", m, "，s值：", s)
	}

	return s
}

// 将2进制字符串转成10进制数值(传统计算方式)
// 样例
// str := "0b10011"
// aa := test_new.Base2To10Custom(str)
// fmt.Println(aa)
// bb := test_new.Base10To2Custom(aa)
// fmt.Println(bb)
// 打印
// 19
// 10011
func Base2To10Custom(str string) int64 {
	//校验
	if str == "" {
		return 0
	}

	//获取字符串中指定字符的前缀(指定字符之前的字符串)，指定字符不存在时返回原字符串
	str = GetPrefixV(str, "0b")

	//2进制
	b := 2
	length := len(str)
	valueLength := length //临时变量，用于记录个十百位的标记

	var sum int64 = 0
	for i := 0; i < length; i++ {

		aa := str[i]     //得到ascii码数字值
		bb := string(aa) //将ascii码数字值转成对应字符串

		f_base := float64(b)
		f_tempFlag := float64(valueLength - 1)
		//fmt.Println("位数：", f_tempFlag)
		value := math.Pow(f_base, f_tempFlag)

		//fmt.Println("value值：", value)

		value = value * StringToFloat64(bb)

		sum += Float64ToInt64(value)

		//位数自减
		valueLength--

		//当位数为0时，停止循环
		if valueLength == 0 {
			break
		}
	}

	//fmt.Println("将2进制转成10进制后：", n)

	return sum
}

// 将8进制字符串转成10进制数值(传统计算方式)
// 样例
// str := "0o177"
// aa := test_new.Base8To10Custom(str)
// fmt.Println(aa)
// bb := test_new.Base10To8Custom(aa)
// fmt.Println(bb)
// 打印
// 127
// 177
func Base8To10Custom(str string) int64 {
	//校验
	if str == "" {
		return 0
	}

	//获取字符串中指定字符的前缀(指定字符之前的字符串)，指定字符不存在时返回原字符串
	str = GetPrefixV(str, "0o")

	//8进制
	b := 8
	length := len(str)
	valueLength := length //临时变量，用于记录个十百位的标记

	var sum int64 = 0
	for i := 0; i < length; i++ {

		aa := str[i]     //得到ascii码数字值
		bb := string(aa) //将ascii码数字值转成对应字符串

		f_base := float64(b)
		f_tempFlag := float64(valueLength - 1)
		//fmt.Println("位数：", f_tempFlag)
		value := math.Pow(f_base, f_tempFlag)

		//fmt.Println("value值：", value)

		value = value * StringToFloat64(bb)

		sum += Float64ToInt64(value)

		//位数自减
		valueLength--

		//当位数为0时，停止循环
		if valueLength == 0 {
			break
		}
	}

	//fmt.Println("将8进制转成10进制后：", n)

	return sum
}

// 将16进制字符串转成10进制数值(传统计算方式)
// 样例
// str := "0xe10a"
// aa := test_new.Base16To10Custom(str)
// fmt.Println(aa)
// bb := test_new.Base10To16Custom(aa)
// fmt.Println(bb)
// 打印
// 57610
// E10A
func Base16To10Custom(str string) int64 {
	//校验
	if str == "" {
		return 0
	}

	//获取字符串中指定字符的前缀(指定字符之前的字符串)，指定字符不存在时返回原字符串
	str = GetPrefixV(str, "0x")

	//16进制
	b := 16
	//str := "1FF"
	length := len(str)
	valueLength := length //临时变量，用于记录个十百位的标记

	var sum int64 = 0
	for i := 0; i < length; i++ {

		aa := str[i]             //得到ascii码数字值
		bb := string(aa)         //将ascii码数字值转成对应字符串
		bb = strings.ToUpper(bb) //将字符串中的字母转成大写
		//fmt.Println("ascii码值：", aa, "转成string后--->", bb)
		//ascii码值： 49 转成string后---> 1
		//ascii码值： 70 转成string后---> F
		//ascii码值： 70 转成string后---> F

		switch bb {
		case "A":
			bb = "10"
		case "B":
			bb = "11"
		case "C":
			bb = "12"
		case "D":
			bb = "13"
		case "E":
			bb = "14"
		case "F":
			bb = "15"
		}

		f_base := float64(b)
		f_tempFlag := float64(valueLength - 1)
		//fmt.Println("位数：", f_tempFlag)
		value := math.Pow(f_base, f_tempFlag)

		//fmt.Println("value值：", value)

		value = value * StringToFloat64(bb)

		sum += Float64ToInt64(value)

		//位数自减
		valueLength--

		//当位数为0时，停止循环
		if valueLength == 0 {
			break
		}
	}

	//fmt.Println("将16进制转成10进制后：", n)

	return sum
}

// 将10进制数值转成62进制字符串(传统计算方式)
// 样例
// aa := test_new.Base10To62Custom(1680000000)
// fmt.Println("10转16后：", aa)
// bb := test_new.Base62To10Custom(aa)
// fmt.Println(n, "--->", aa, "--->", bb)
// 打印
// 1680000000 ---> 1ph6kC ---> 1680000000
func Base10To62Custom(n int64) string {
	if n < 1 {
		return "0"
	}

	//62进制
	var b int64 = 62

	s := ""
	for i := n; i > 0; i = i / b {
		m := i % b
		if m > 9 && m < b { //如果余数是10到15之间的数时
			temp := ""
			switch m {
			case 10:
				temp = "A"
			case 11:
				temp = "B"
			case 12:
				temp = "C"
			case 13:
				temp = "D"
			case 14:
				temp = "E"
			case 15:
				temp = "F"
			case 16:
				temp = "G"
			case 17:
				temp = "H"
			case 18:
				temp = "I"
			case 19:
				temp = "J"
			case 20:
				temp = "K"
			case 21:
				temp = "L"
			case 22:
				temp = "M"
			case 23:
				temp = "N"
			case 24:
				temp = "O"
			case 25:
				temp = "P"
			case 26:
				temp = "Q"
			case 27:
				temp = "R"
			case 28:
				temp = "S"
			case 29:
				temp = "T"
			case 30:
				temp = "U"
			case 31:
				temp = "V"
			case 32:
				temp = "W"
			case 33:
				temp = "X"
			case 34:
				temp = "Y"
			case 35:
				temp = "Z"
			case 36:
				temp = "a"
			case 37:
				temp = "b"
			case 38:
				temp = "c"
			case 39:
				temp = "d"
			case 40:
				temp = "e"
			case 41:
				temp = "f"
			case 42:
				temp = "g"
			case 43:
				temp = "h"
			case 44:
				temp = "i"
			case 45:
				temp = "j"
			case 46:
				temp = "k"
			case 47:
				temp = "l"
			case 48:
				temp = "m"
			case 49:
				temp = "n"
			case 50:
				temp = "o"
			case 51:
				temp = "p"
			case 52:
				temp = "q"
			case 53:
				temp = "r"
			case 54:
				temp = "s"
			case 55:
				temp = "t"
			case 56:
				temp = "u"
			case 57:
				temp = "v"
			case 58:
				temp = "w"
			case 59:
				temp = "x"
			case 60:
				temp = "y"
			case 61:
				temp = "z"
			}

			//s = fmt.Sprintf("%v%v", temp, s) //将000111这样的值一个一个字符串拼接给临时变量s
			s = temp + s //将000111这样的值一个一个字符串拼接给临时变量s
			//fmt.Println("循环内--->i值：", i, "，m值：", m, "，s值：", s)
			continue //跳出，下面的代码不执行，很关键，一定要写才行。
		}

		//s = fmt.Sprintf("%v%v", m, s) //将000111这样的值一个一个字符串拼接给临时变量s
		//s = strconv.Itoa(m) + s //将000111这样的值一个一个字符串拼接给临时变量s
		s = Int64ToString(m) + s //将000111这样的值一个一个字符串拼接给临时变量s
		//fmt.Println("循环内--->i值：", i, "，m值：", m, "，s值：", s)
	}

	return s
}

// 将62进制字符串转成10进制数值(传统计算方式)
// 样例
// aa := test_new.Base10To62Custom(1680000000)
// fmt.Println("10转16后：", aa)
// bb := test_new.Base62To10Custom(aa)
// fmt.Println(n, "--->", aa, "--->", bb)
// 打印
// 1680000000 ---> 1ph6kC ---> 1680000000
func Base62To10Custom(str string) int64 {
	//校验
	if str == "" {
		return 0
	}

	//62进制
	var b int64 = 62
	//str := "1FF"
	length := len(str)
	valueLength := length //临时变量，用于记录个十百位的标记

	var sum int64 = 0
	for i := 0; i < length; i++ {

		aa := str[i]     //得到ascii码数字值
		bb := string(aa) //将ascii码数字值转成对应字符串
		//bb = strings.ToUpper(bb) //将字符串中的字母转成大写
		//注意，这里不能将小写字母转大写字母，因为下面有小写字母的判断。因为是62进制，使用了所有大小写字母，大写字母与小写字母的所对应的值是不同的，
		//这里不能和16进制相比，因为16进制只使用了a到f。
		//fmt.Println("ascii码值：", aa, "转成string后--->", bb)
		//ascii码值： 49 转成string后---> 1
		//ascii码值： 70 转成string后---> F
		//ascii码值： 70 转成string后---> F

		switch bb {
		case "A":
			bb = "10"
		case "B":
			bb = "11"
		case "C":
			bb = "12"
		case "D":
			bb = "13"
		case "E":
			bb = "14"
		case "F":
			bb = "15"
		case "G":
			bb = "16"
		case "H":
			bb = "17"
		case "I":
			bb = "18"
		case "J":
			bb = "19"
		case "K":
			bb = "20"
		case "L":
			bb = "21"
		case "M":
			bb = "22"
		case "N":
			bb = "23"
		case "O":
			bb = "24"
		case "P":
			bb = "25"
		case "Q":
			bb = "26"
		case "R":
			bb = "27"
		case "S":
			bb = "28"
		case "T":
			bb = "29"
		case "U":
			bb = "30"
		case "V":
			bb = "31"
		case "W":
			bb = "32"
		case "X":
			bb = "33"
		case "Y":
			bb = "34"
		case "Z":
			bb = "35"
		case "a":
			bb = "36"
		case "b":
			bb = "37"
		case "c":
			bb = "38"
		case "d":
			bb = "39"
		case "e":
			bb = "40"
		case "f":
			bb = "41"
		case "g":
			bb = "42"
		case "h":
			bb = "43"
		case "i":
			bb = "44"
		case "j":
			bb = "45"
		case "k":
			bb = "46"
		case "l":
			bb = "47"
		case "m":
			bb = "48"
		case "n":
			bb = "49"
		case "o":
			bb = "50"
		case "p":
			bb = "51"
		case "q":
			bb = "52"
		case "r":
			bb = "53"
		case "s":
			bb = "54"
		case "t":
			bb = "55"
		case "u":
			bb = "56"
		case "v":
			bb = "57"
		case "w":
			bb = "58"
		case "x":
			bb = "59"
		case "y":
			bb = "60"
		case "z":
			bb = "61"
		}

		f_base := float64(b)
		f_tempFlag := float64(valueLength - 1)
		//fmt.Println("位数：", f_tempFlag)
		value := math.Pow(f_base, f_tempFlag)

		//fmt.Println("value值：", value)

		value = value * StringToFloat64(bb)

		sum += Float64ToInt64(value)

		//位数自减
		valueLength--

		//当位数为0时，停止循环
		if valueLength == 0 {
			break
		}
	}

	//fmt.Println("将62进制转成10进制后：", n)

	return sum
}

// ====================================传统计算方式封装的函数========================================
