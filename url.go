package liuyang

import (
	"net/url"
)

// 将字符串进行urlencode编码
// 样例
// str := "1680000058_b_刘阳_/a/b/c?id=8&tid=3"
// aa := test_new.UrlEncode(str)
// fmt.Println(aa)
// bb := test_new.UrlDecode(aa)
// fmt.Println(bb)
// 打印
// 1680000058_b_%E5%88%98%E9%98%B3_%2Fa%2Fb%2Fc%3Fid%3D8%26tid%3D3
// 1680000058_b_刘阳_/a/b/c?id=8&tid=3
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

// 将urlencode字符串进行解码
func UrlDecode(s string) string {
	res, err := url.QueryUnescape(s)
	if err != nil {
		return ""
	}
	return res
}
