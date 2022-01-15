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
//进制转换，2转10，int类型，base为0时，代表参数可以传以0b、0o、0x开头的字符串，如果base大于0就不允许传特殊字符。
//	ab2 := test.Base0To10Int("0b10110") //22
//	ab8 := test.Base0To10Int("0o26")    //22
//	ab16 := test.Base0To10Int("0x16")   //22
func Base0To10Int(s string) int64 {
	n, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0
	}

	return n
}

//进制转换，2转10，int类型
//	aa2 := test.Base2To10UInt("10110") //22
//	aa8 := test.Base8To10UInt("26")    //22
//	aa16 := test.Base16To10UInt("16")  //22
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

//----------------打印[]byte数组，里面包含10进制数，得到2、8、16进制数，参数默认0包含0x和补0----------------
//打印[]byte数组，里面包含的是10进制数，得到0b开头的2进制数并以分号连接的字符串
func PrintBase10To2(byteArr []byte, flag uint8) string {
	var str = ""       //总总符串，默认空
	var prefix_0 = ""  //补字符0，默认空
	var prefix_0b = "" //补字符0b,默认空
	if flag == 0 {
		prefix_0 = "0"
		prefix_0b = "0b"
	} else if flag == 1 {
		prefix_0 = "0"
		prefix_0b = ""
	} else if flag == 2 {
		prefix_0 = ""
		prefix_0b = "0b"
	} else if flag == 3 {
		prefix_0 = ""
		prefix_0b = ""
	}
	for key, value := range byteArr {
		//将10进制数转成16进制字符串
		cardData := BaseUInt10To2(uint64(value))
		//如果长度为1，也就是0到9时，补0，这样方便看，不然会是0x1,0x2，远没有0x01,0x02好看。
		if len(cardData) == 1 {
			cardData = prefix_0 + cardData
		}
		if key < len(byteArr)-1 {
			str += prefix_0b + cardData + ","
		} else {
			str += prefix_0b + cardData
		}
	}

	return str
}

//打印[]byte数组，里面包含的是10进制数，得到0o开头的8进制数并以分号连接的字符串
func PrintBase10To8(byteArr []byte, flag uint8) string {
	var str = ""       //总总符串，默认空
	var prefix_0 = ""  //补字符0，默认空
	var prefix_0b = "" //补字符0b,默认空
	if flag == 0 {
		prefix_0 = "0"
		prefix_0b = "0o"
	} else if flag == 1 {
		prefix_0 = "0"
		prefix_0b = ""
	} else if flag == 2 {
		prefix_0 = ""
		prefix_0b = "0o"
	} else if flag == 3 {
		prefix_0 = ""
		prefix_0b = ""
	}
	for key, value := range byteArr {
		//将10进制数转成16进制字符串
		cardData := BaseUInt10To8(uint64(value))
		//如果长度为1，也就是0到9时，补0，这样方便看，不然会是0x1,0x2，远没有0x01,0x02好看。
		if len(cardData) == 1 {
			cardData = prefix_0 + cardData
		}
		if key < len(byteArr)-1 {
			str += prefix_0b + cardData + ","
		} else {
			str += prefix_0b + cardData
		}
	}

	return str
}

//打印[]byte数组，里面包含的是10进制数，得到0x开头的16进制数并以分号连接的字符串
func PrintBase10To16(byteArr []byte, flag uint8) string {
	var str = ""       //总总符串，默认空
	var prefix_0 = ""  //补字符0，默认空
	var prefix_0b = "" //补字符0b,默认空
	if flag == 0 {
		prefix_0 = "0"
		prefix_0b = "0x"
	} else if flag == 1 {
		prefix_0 = "0"
		prefix_0b = ""
	} else if flag == 2 {
		prefix_0 = ""
		prefix_0b = "0x"
	} else if flag == 3 {
		prefix_0 = ""
		prefix_0b = ""
	}
	for key, value := range byteArr {
		//将10进制数转成16进制字符串
		cardData := BaseUInt10To16(uint64(value))
		//如果长度为1，也就是0到9时，补0，这样方便看，不然会是0x1,0x2，远没有0x01,0x02好看。
		if len(cardData) == 1 {
			cardData = prefix_0 + cardData
		}
		if key < len(byteArr)-1 {
			str += prefix_0b + cardData + ","
		} else {
			str += prefix_0b + cardData
		}
	}

	return str
}

//----------------打印[]byte数组，里面包含10进制数，得到2、8、16进制数，参数默认0包含0x和补0----------------
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
