package liuyang

import "regexp"

// 检查用户名只能输入字母或数字或下划线(5至20位)，以字母开头。
func CheckUsername(str string) bool {
	expreg := "^[a-zA-Z]\\w{4,19}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "123fa2fs314569_87as3"
	istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	return istrue
}

// 检查密码，6-20位。
func CheckPassword(str string) bool {
	expreg := "^[!@#$%^&*\\w]{6,20}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查手机号
func CheckMobile(str string) bool {
	expreg := "^[1][3-9]\\d{9}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查邮箱
func CheckEmail(str string) bool {
	expreg := "^\\w{1,30}@\\w{1,20}(\\.[a-zA-Z]{1,10}){1,2}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 只能输入汉字、字母、数字、下划线(1至16个)
func CheckTitle(str string) bool {
	expreg := "^[\u4e00-\u9fa5\\w]{1,16}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查身份证码码(18位身份证号码，年月日严格校验)
// 样例1
// id := "230121008611133619"
// ok := test_new.Check_IDCard(id)
// 打印
// false
// 样例2
// id := "230121198611133619"
// ok := test_new.Check_IDCard(id)
// 打印
// true
// 样例2
// id := "230121998611133619"
// ok := test_new.Check_IDCard(id)
// 打印
// true
func CheckIDCard(str string) bool {
	expreg := "^[1-9]\\d{5}(([1][9]\\d{2})|([2-9]\\d{3}))([0][1-9]|[1][012])([0][1-9]|[12][0-9]|[3][01])(\\d{4}|\\d{3}[X])$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查座机号码
func CheckTelephone(str string) bool {
	expreg := "^([0]\\d{2,3}-\\d{6,8}(-\\d{1,6})?)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查网址是否合法
func CheckUrl(str string) bool {
	expreg := "^(\\w+)[:]\\/\\/([a-zA-Z0-9][a-zA-Z0-9-]+)[.]([a-zA-Z0-9][a-zA-Z0-9-]+)+([:]\\d*)?([^:?]*)?([^#/]*)?(#.*)?"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "http://aDmin.fc114.COM.cn:12345/site"
	//str := "http://liuyang 15067116661"
	//arr := reg.FindAllString(str,1)
	//fmt.Println(arr)
	istrue := reg.MatchString(str)

	return istrue
}

// 检查金额，只能输入大于0并且小于1000亿的正整数或小数(精确0至2位)，刘阳推荐
func CheckMoney(str string) bool {
	expreg := "^((0\\.(0[1-9]|[1-9]\\d?))|([1-9]\\d{0,11}(\\.\\d{1,2})?))$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查(正负)金额，只能输入大于0并且小于1000亿的正整数或小数(精确0至2位)，刘阳推荐
func CheckSMoney(str string) bool {
	expreg := "^(\\+|\\-)?((0\\.(0[1-9]|[1-9]\\d?))|([1-9]\\d{0,11}(\\.\\d{1,2})?))$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查小数，可输入0.00，或输入小于1000亿的正整数或小数(精确0至2位)，刘阳推荐
func CheckDecimal(str string) bool {
	expreg := "^((0|[1-9]\\d{0,11})(\\.\\d{1,2})?)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查(正负)小数，可输入0.00，或输入小于1000亿的正整数或小数(精确0至2位)，刘阳推荐
func CheckSDecimal(str string) bool {
	expreg := "^((\\+|\\-)?(0|[1-9]\\d{0,11})(\\.\\d{1,2})?)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查小数(6位)，只能输入大于0并且小于1000亿的正整数或小数(精确0至6位)
func CheckDecimal6(str string) bool {
	expreg := "^((0\\.(\\d{0,1}[1-9]\\d{0,4}|\\d{0,2}[1-9]\\d{0,3}|\\d{0,3}[1-9]\\d{0,2}|\\d{0,4}[1-9]\\d{0,1}|\\d{0,5}[1-9]|[1-9]\\d{0,5}))|([1-9]\\d{0,11}(\\.\\d{1,6})?))$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查(正负)小数(6位)，只能输入大于0并且小于1000亿的正整数或小数(精确0至6位)
func CheckSDecimal6(str string) bool {
	expreg := "^(\\+|\\-)?((0\\.(\\d{0,1}[1-9]\\d{0,4}|\\d{0,2}[1-9]\\d{0,3}|\\d{0,3}[1-9]\\d{0,2}|\\d{0,4}[1-9]\\d{0,1}|\\d{0,5}[1-9]|[1-9]\\d{0,5}))|([1-9]\\d{0,11}(\\.\\d{1,6})?))$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查数字，可输入0，或小于uint64最大值的正整数
func Check_0_20(str string) bool {
	expreg := "^(0|[1-9]\\d{0,19})$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查数字，只能输入大于或等于1并且小于uint64最大值的正整数
func Check_1_20(str string) bool {
	expreg := "^([1-9]\\d{0,19})$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查纯数字，不限制个数
func CheckNumber(str string) bool {
	expreg := "^(\\d+)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查(正负)纯数字，不限制个数
func CheckSNumber(str string) bool {
	expreg := "^((\\+|\\-)?\\d+)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查用户id，以大于0值开头的数字且小于19位为合法。
func CheckUid(str string) bool {
	expreg := "^[1-9]\\d{0,18}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查md5格式是否合法，判定长度为固定32位的数字和字母
func CheckMD5(str string) bool {
	expreg := "^[0-9a-zA-Z]{32}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "123fa2fs314569_87as3"
	istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	return istrue
}

// 检查token格式是否合法，判定为30至100之间的数字和字母
func CheckToken(str string) bool {
	expreg := "^[0-9a-zA-Z]{30,100}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "123fa2fs314569_87as3"
	istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	return istrue
}

// 检查jwt格式是否合法，判定为50至300之间的任意字符(不包括',")
func CheckJwt(str string) bool {
	expreg := "^[^'\"]{50,300}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "123fa2fs314569_87as3"
	istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	return istrue
}

// 检查设备id格式是否合法，判定为6至200之间的任意字符(不包括',")
func CheckDevid(str string) bool {
	expreg := "^[^'\"]{6,200}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	//str := "123fa2fs314569_87as3"
	istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	return istrue
}

// 检查版本号，如：9.0、9.1、9.10、9.10.1、9.10.10.1、10.10.0.0.0.1
func CheckVersion(str string) bool {
	expreg := "^\\d+(.\\d+)*$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查坐标，如：1.2546,34.5456456
func CheckPosition(str string) bool {
	expreg := "^-?\\d+(\\.\\d+)?,-?\\d+(\\.\\d+)?$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查字符串是否为单个或以英文逗号连接的多个字符串(实战风格多选或图片地址)
// 样例
// 1.2546,-34.5456456,2sflasdf,asdf,asdfasdf
func CheckStrStr(str string) bool {
	expreg := "^.+(,.+)*$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查字符串是否为多个整型数字字符串(实战为门店或集团门店传值，不允许有0值)
// 样例
// id := "11.ff,,abc,def,ghi"
// ok := test_new.CheckNumberNumber(id)
// fmt.Println(ok)
// 打印
// false
// 样例2
// id := "30230,10,20,3538,6021"
// ok := test_new.CheckNumberNumber(id)
// fmt.Println(ok)
// 打印
// true
func CheckNumberNumber(str string) bool {
	expreg := "^[1-9]\\d*(,[1-9]\\d*)*$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查字符串是否为json字符串
// 样例1
// str := `liuyang`
// ok := test_new.CheckJson(str)
// 打印
// false
// 样例2
// str := `{"name":"liuyang","a":[]}`
// ok := test_new.CheckJson(str)
// 打印
// true
// 样例3
// str := `[{"name":"liuyang","a":[]}]`
// ok := test_new.CheckJson(str)
// 打印
// true
// 注意：两个反斜杠是go语言特殊处理，在php中只用一个斜杠用于转义，而在go语言中就得用两个斜杠才行。
// php中样例：$reg = "/^((\[.*])|(\{.*})|(\[{.*}]))$/";
func CheckJson(str string) bool {
	//只允许输入json字符串：[]、{}、[{}]
	//注意：不要校验里面的冒号，因为像[2,3,5]这样的普通数组中是没有冒号的
	expreg := "^((\\[.*])|(\\{.*})|(\\[{.*}]))$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

// 检查json字符串是否已加上转义符
// 样例1
// str := `[{"name":"liuyang","a":[]}]`
// ok := test_new.CheckJsonEscape(str)
// fmt.Println(ok)
// 打印
// false
// 样例2
// str := `{\"id\":25,\"title\":\"\\u5218\\u9633\",\"mob\":\"15067116661\",\"flag\":1}`
// ok := test_new.CheckJsonEscape(str)
// fmt.Println(ok)
// 打印
// true
// 样例3
// str := `[{\"id\":108,\"title\":\"\\u6768\\u6587\\u79c0\",\"mob\":\"15700091969\",\"flag\": 1 }]`
// ok := test_new.CheckJsonEscape(str)
// fmt.Println(ok)
// 打印
// true
// 注意：两个反斜杠是go语言特殊处理，在php中用了三个斜杠用于转义，因为有一个要将php固有的两个//转义。而在go语言中用两个斜杠就行。
// php中样例：$reg = '/\\\"/';
func CheckJsonEscape(str string) bool {
	//检查双引号前面是否已加上转义符(斜杠)，如：" 转义后 \"
	expreg := `\\"`
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}
