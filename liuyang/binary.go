package liuyang

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//二进制测试
func Binary_test()  {

	aa := 0x1000
	fmt.Println("0x0010十进制数是：",aa)
	//调用方法：
	result2 := DecToBin(int64(aa))
	fmt.Println("调用二进制方法后：",result2)

	resulta2, rerra2 := strconv.ParseInt(result2, 2, 64)
	fmt.Println("二进制转十进制后：",rerra2,resulta2)

	result8 := DecToOct(int64(aa))
	fmt.Println("调用八进制方法后：",result8)

	resulta8, rerra8 := strconv.ParseInt(result8, 8, 64)
	fmt.Println("八进制转十进制后：",rerra8,resulta8)

	result16 := DecToHex(int64(aa))
	fmt.Println("调用16进制方法后：",result16)

	resulta16, rerra16 := strconv.ParseInt(result16, 16, 64)
	fmt.Println("十六进制转十进制后：",rerra16,resulta16)

	fmt.Println("=======================================")
	fmt.Println("2转8：",BinToOct(result2))//2转8： 70
	fmt.Println("2转16：",BinToHex(result2))//2转16： 38
	fmt.Println("8转2：",OctToBin(result8))//8转2： 111000
	fmt.Println("8转16：",OctToHex(result8))//8转16： 38
	fmt.Println("16转2：",HexToBin(result16))//16转2： 111000
	fmt.Println("16转8：",HexToOct(result16))//16转8： 70

}


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
		s = fmt.Sprintf("%v%v", m, s)//将000111这样的值一个一个字符串拼接给临时变量s
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
		s = fmt.Sprintf("%v%v", m, s)//将000111这样的值一个一个字符串拼接给临时变量s
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
		if m > 9 && m < 16 {//如果余数是10到15之间的数时
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

			s = fmt.Sprintf("%v%v", temp, s)//将000111这样的值一个一个字符串拼接给临时变量s
			continue	//跳出，下面的代码不执行，很关键，一定要写才行。
		}

		s = fmt.Sprintf("%v%v", m, s)//将000111这样的值一个一个字符串拼接给临时变量s
	}
	return s
}

//将2进制转8进制
func BinToOct(binStr string) string {
	//先将2进制转成10进制
	dec, _ := strconv.ParseInt(binStr, 2, 64)
	//再将10进制转8进制
	octStr := DecToOct(dec)

	return octStr
}

//将2进制转16进制
func BinToHex(binStr string) string {
	//先将2进制转成10进制
	dec, _ := strconv.ParseInt(binStr, 2, 64)
	//再将10进制转16进制
	hexStr := DecToHex(dec)

	return hexStr
}


//将8进制转2进制
func OctToBin(octStr string) string {
	//先将8进制转成10进制
	dec, _ := strconv.ParseInt(octStr, 8, 64)
	//再将10进制转2进制
	binStr := DecToBin(dec)

	return binStr
}

//将8进制转16进制
func OctToHex(octStr string) string {
	//先将8进制转成10进制
	dec, _ := strconv.ParseInt(octStr, 8, 64)
	//再将10进制转16进制
	hexStr := DecToHex(dec)

	return hexStr
}


//将16进制转2进制
func HexToBin(hexStr string) string {
	//先将16进制转成10进制
	dec, _ := strconv.ParseInt(hexStr, 16, 64)
	//再将10进制转2进制
	binStr := DecToBin(dec)

	return binStr
}

//将16进制转8进制
func HexToOct(hexStr string) string {
	//先将16进制转成10进制
	dec, _ := strconv.ParseInt(hexStr, 16, 64)
	//再将10进制转8进制
	octStr := DecToOct(dec)

	return octStr
}


/////////////////////////////////////////////////////////////////////////////////
//网上原文代码：
// Decimal to binary
func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// Decimal to octal
func DecOct(d int64) int64 {
	if d == 0 {
		return 0
	}
	if d < 0 {
		log.Println("Decimal to octal error: the argument must be greater than zero.")
		return -1
	}
	s := ""
	for q := d; q > 0; q = q / 8 {
		m := q % 8
		s = fmt.Sprintf("%v%v", m, s)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Println("Decimal to octal error:", err.Error())
		return -1
	}
	return int64(n)
}

// Decimal to hexadecimal
func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// Binary to decimal
func BinDec(b string) (n int64) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d)
}

// Octal to decimal
func OctDec(o int64) (n int64) {
	s := strings.Split(strconv.Itoa(int(o)), "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Octal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(8, float64(l-i-1))
	}
	return int64(d)
}

// Hexadecimal to decimal
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

// Octal to binary
func OctBin(o int64) string {
	d := OctDec(o)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// Hexadecimal to binary
func HexBin(h string) string {
	d := HexDec(h)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// Binary to octal
func BinOct(b string) int64 {
	d := BinDec(b)
	if d == -1 {
		return -1
	}
	return DecOct(d)
}

// Binary to hexadecimal
func BinHex(b string) string {
	d := BinDec(b)
	if d == -1 {
		return ""
	}
	return DecHex(d)
}


//该方法是刘阳写的最笨重的方法，不推荐使用。
func Test()  {
	a := 56
	temp := 0

	arr := make([]int,0)

	//该方法是刘阳写的最笨重的方法，不推荐使用。
	for {
		if a >= 2 {
			temp = a % 2				//余数有时为0，有时为1
			arr = append(arr, temp)		//将余数插入数组中
			a /= 2						//除以2后的下一级数，循环往复->取余->除2
			fmt.Println("a的值是；",a)
		} else {
			arr = append(arr, a)		//将最后的余数插入数组中
			break
		}
	}
	fmt.Println("二进制的值是：",arr)


	//倒序排序
	list := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		list = append(list, arr[i])
	}

	fmt.Println("重新排序后：", list)
}
