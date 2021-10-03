package liuyang

import (
	"math"
	"strconv"
	"strings"
	"unsafe"
)

//比较两个英文字符串的大小
//字符串1大于字符串2时，返回：1
//字符串2小于字符串2时，返回：-1
//字符串1和字符串2相等时，返回：0
//虽然和系统方法功能一样，但还是建议使用系统方法：strings.Compare()
func StringCompareInt(str1, str2 string) int {
	lenS1 := len(str1)
	lenS2 := len(str2)
	for key,_ := range str1 {
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
				return +1
			}
			continue//此处的continue是使用的最佳时机
		} else if str1[key] > str2[key] {
			//fmt.Println("这是字母1大于字母2")
			return +1
		} else {
			//fmt.Println("这是字母2大于字母1")
			return -1
		}
		//fmt.Println(str1[key],"==>",str2[key])
	}

	return 0
}

//判断字符串是否为空，如果有空格的话，也视为空
func StringIsEmpty(str string) bool {

	//首先过滤掉字符串中的首尾空白字符
	tempStr := StringTrim(str)
	//再判断其长度是否大于0，如果大于0，说明不为空
	if len(tempStr) > 0 {
		return false
	}

	return true
}
//过滤掉字符串中的首尾空白字符，包括空格、制表符、换页符等等。
func StringTrim(str string) string {
	//过滤掉字符串中的首尾空白字符，包括空格、制表符、换页符等等。
	tempStr := strings.Trim(str," \f\n\r\t\v")
	return tempStr
}

//得到一个英文字符串的ASCII码值的总和
func StringAscllSum(str string) int {
	var sum int = 0
	//fmt.Println(str)//liu yang
	for i := 0; i < len(str); i++ {
		//fmt.Print(str[i],"\t")//108	105	117	32	121	97	110	103
		sum += int(str[i])
	}
	//fmt.Println(sum)//793

	return sum
}

//string类型转int8
func StringToInt8(s string) int8 {
	n,err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt8 {
		return 0
	}

	return int8(n)
}

//string类型转uint8
func StringToUInt8(s string) uint8 {
	n,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint8 {
		return 0
	}

	return uint8(n)
}

//string类型转int16
func StringToInt16(s string) int16 {
	n,err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt16 {
		return 0
	}

	return int16(n)
}

//string类型转uint16
func StringToUInt16(s string) uint16 {
	n,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint16 {
		return 0
	}

	return uint16(n)
}


//string类型转int32
func StringToInt32(s string) int32{
	n,err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt32 {
		return 0
	}

	return int32(n)
}

//string类型转uint32
func StringToUInt32(s string) uint32 {
	n,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxUint32 {
		return 0
	}

	return uint32(n)
}

//string类型转int64
func StringToInt64(s string) int64 {
	n,err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if n > math.MaxInt64 {
		return 0
	}

	return n
}

//string类型转uint64
func StringToUInt64(s string) uint64 {
	n,err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}

	return n
}

//string类型转float32，虽然返回的是float64类型，但是以float32为精度截取的
func StringToFloat32(s string) float32 {
	//str := "123456.0123456789653"
	//fmt.Println(a1)//123456.015625
	ff,err := strconv.ParseFloat(s, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}
	//如果转后后的数值大于当前类型的最大值时，返回0
	if ff > math.MaxFloat32 {
		return 0
	}

	return float32(ff)
}

//string转float64
func StringToFloat64(s string) float64 {
	//str := "123456.0123456789653"
	//fmt.Println(a2)//123456.01234567896
	//值的注意：string转float时，只能转6位数以下的值，超过6位就会变成科学计数法，如下：
	//str := "123456"
	//fmt.Println(a2)//123456
	//str := "1234567"
	//fmt.Println(a2)//1.234567e+06
	ff,err := strconv.ParseFloat(s, 64)
	if err != nil {
		//字符串中的数字大于64位最大值19位时的报错
		return 0//这里不用返回该类型的最大值，返回0就代表转换失败
	}

	return ff
}

//新增
//string转byte的高级写法
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

//新增
//byte转string的高级写法，bytes2string将字节片转换为字符串，无需内存分配
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
