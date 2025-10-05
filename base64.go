package liuyang

import "encoding/base64"

// 将字符串进行base64编码
// 样例
// str := "liuyang"
// aa := test_new.Base64Encode(str)
// fmt.Println(aa)
// bb := test_new.Base64Decode(aa)
// fmt.Println(bb)
// 打印
// bGl1eWFuZw==
// liuyang
func Base64Encode(s string) string {
	data := []byte(s)
	return base64.StdEncoding.EncodeToString(data)
}

// 将base64字符串进行解码
func Base64Decode(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(data)
}
