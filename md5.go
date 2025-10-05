package liuyang

import (
	"crypto/md5"
	"fmt"
)

// 将字符串进行md5加密
// 样例
// str := "123456"
// md5Str := test_new.Md5(str)
// fmt.Println(md5Str)
// 打印
// e10adc3949ba59abbe56e057f20f883e
func Md5(str string) string {
	//将string数据强转[]byte类型
	data := []byte(str)
	byteArr16 := md5.Sum(data)
	//得到[16]byte数组，不是切片，不能用ByteToString转，只能用Sprintf转

	tempStr := fmt.Sprintf("%x", byteArr16)

	return tempStr
}
