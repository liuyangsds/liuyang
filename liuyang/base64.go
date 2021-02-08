package liuyang

import (
	"encoding/base64"
	"fmt"
)

//base64加解密测试
func Base64_test()  {
	//base64加密
	byteStr := []byte("刘阳是好人")
	str := base64.StdEncoding.EncodeToString(byteStr)
	fmt.Println("base64加密后：",str)//5YiY6Ziz5piv5aW95Lq6

	//base64解密
	deByte,deErr := base64.StdEncoding.DecodeString(str)
	if deErr != nil {
		fmt.Println(deErr)
	}
	fmt.Println("base64解密后：",string(deByte))//刘阳是好人

}

//加密
func Base64_encode(src []byte) string {

	return base64.StdEncoding.EncodeToString(src)
}

//解密
func Base64_decode(s string) ([]byte, error) {

	return base64.StdEncoding.DecodeString(s)
}