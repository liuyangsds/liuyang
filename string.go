package liuyang

import (
	"strings"
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
func StringWordSum(str string) int {
	var sum int = 0
	//fmt.Println(str)//liu yang
	for i := 0; i < len(str); i++ {
		//fmt.Print(str[i],"\t")//108	105	117	32	121	97	110	103
		sum += int(str[i])
	}
	//fmt.Println(sum)//793

	return sum
}
