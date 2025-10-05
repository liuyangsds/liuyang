package liuyang

import (
	"crypto/sha1"
	"fmt"
)

// 将字符串进行sha1加密
// 样例
// str := "123456"
// tempStr := test_new.Sha1(str)
// fmt.Println(tempStr)
// 打印
// 7c4a8d09ca3762af61e59520943dc26494f8941b
func Sha1(str string) string {
	//将string数据强转[]byte类型
	data := []byte(str)
	byteArr20 := sha1.Sum(data)
	//得到[20]byte数组，不是切片，不能用ByteToString转，只能用Sprintf转
	tempStr := fmt.Sprintf("%x", byteArr20)

	return tempStr
}
