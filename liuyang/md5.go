package liuyang

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

//md5测试
func MD5_test()  {
	str := "123456"

	//方式1
	data := []byte(str)
	md5x := md5.Sum(data)
	md5str := fmt.Sprintf("%x",md5x)
	fmt.Println(md5str)
	fmt.Println("----------------------------")

	//方式2
	w := md5.New()
	io.WriteString(w,str)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(md5str2)
	fmt.Println("----------------------------")

	//方式3
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))


	//打印类型：
	//tt := reflect.TypeOf(md)		//第一种，反射
	//fmt.Println(tt)
	//fmt.Printf("%T",md)		//第二种：直接打印
}

//将字符串转成md5字符串，推荐
func MD5(str string) string {
	data := []byte(str)
	md5x := md5.Sum(data)
	md5str := fmt.Sprintf("%x",md5x)

	return md5str
}

