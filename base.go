package liuyang

import (
	"fmt"
	"log"
	"strconv"
)

//====================================自定义封装进制函数========================================
//----------------10进制转成2、8、16进制----------------
//进制转换，10转2，int类型
func BaseInt10To2(n int64) string {
	return strconv.FormatInt(n, 2)
}

//进制转换，10转2，uint类型
func BaseUInt10To2(n uint64) string {
	return strconv.FormatUint(n, 2)
}

//进制转换，10转8，int类型
func BaseInt10To8(n int64) string {
	return strconv.FormatInt(n, 8)
}

//进制转换，10转8，uint类型
func BaseUInt10To8(n uint64) string {
	return strconv.FormatUint(n, 8)
}

//进制转换，10转16，int类型
func BaseInt10To16(n int64) string {
	return strconv.FormatInt(n, 16)
}

//进制转换，10转2，uint类型
func BaseUInt10To16(n uint64) string {
	return strconv.FormatUint(n, 16)
}

//----------------10进制转成2、8、16进制----------------

//----------------2、8、16进制转成10进制----------------
//进制转换，2转10，int类型
func Base2To10Int(s string) int64 {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，2转10，uint类型
func Base2To10UInt(s string) uint64 {
	n, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，8转10，int类型
func Base8To10Int(s string) int64 {
	n, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，8转10，uint类型
func Base8To10UInt(s string) uint64 {
	n, err := strconv.ParseUint(s, 8, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，16转10，int类型
func Base16To10Int(s string) int64 {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，16转10，uint类型
func Base16To10UInt(s string) uint64 {
	n, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}

	return n
}

//----------------2、8、16进制转成10进制----------------
//====================================自定义封装进制函数========================================

//====================================进制学习专用函数========================================
//将10进制转2进制后得到字符串，刘阳推荐
func DecToBin(n int64) string {
	if n < 0 {
		log.Println("十进制到二进制错误：参数必须大于零")
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

//将10进制转8进制后得到字符串，刘阳推荐
func DecToOct(n int64) string {
	if n < 0 {
		log.Println("10进制到8进制错误：参数必须大于零")
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

//将10进制转16进制后得到字符串，刘阳推荐
func DecToHex(n int64) string {
	if n < 0 {
		log.Println("10进制到16进制错误：参数必须大于零")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 { //如果余数是10到15之间的数时
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

			s = fmt.Sprintf("%v%v", temp, s) //将000111这样的值一个一个字符串拼接给临时变量s
			continue                         //跳出，下面的代码不执行，很关键，一定要写才行。
		}

		s = fmt.Sprintf("%v%v", m, s) //将000111这样的值一个一个字符串拼接给临时变量s
	}
	return s
}

//====================================进制学习专用函数========================================
