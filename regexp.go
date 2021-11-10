package liuyang


import (
	"fmt"
	"regexp"
)

//正则测试
func Check_test()  {
	//简单的检测可以直接用regexp.MatchString函数。
	//而如果想要做更复杂的操作，如：find(js中为search)、replace、split等，就要用regexp.Compile函数做声明并返回指针。
	//这样以返回的正则结构体的指针对象调用其内部方法find()就可以了。
	/*
		如果使用byte，其步骤为：
		1，将字符串类型转换成byte类型。
		2，开始正则匹配。
		3，将返回结果byte类型转换成string类型。

		如果使用string，其步骤为：
		1，获取返回结果即可。

		所以，为了省事，还是使用string类型比较好。
	*/

	//常规写法：
	//bytestr := []byte("123刘阳")
	//rega, err := regexp.Compile("[a-z0-9]")
	//str := rega.Match(bytestr)//bool

	//简化写法：
	//bytelist := []byte("123刘阳")
	//retrue,reerr := regexp.Match("[a-z0-9]",bytelist)//bool,error


	//正则测试：
	//用户名只能输入字母或数字或下划线(5至20位)
	//expreg := "^\\w{5,20}$"
	//reg, _ := regexp.Compile(expreg)
	//str := "123fa2fs314569_87as3"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//用户名只能输入字母或数字或下划线(5至20位)

	//aa,bb := Check_username(str)
	//fmt.Println(bb)
	//fmt.Println(aa)

	//密码可以用特殊字符、大小写字母、数字、下划线。(6至20位)。
	//expreg := "^[!@#$%^&*\\w]{6,20}$"
	//reg, _ := regexp.Compile(expreg)
	//str := "22!@#$%^&*569_87as3a"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//密码可以用特殊字符、大小写字母、数字、下划线。(6至20位)。

	//aa,bb := Check_password(str)
	//fmt.Println(bb)
	//fmt.Println(aa)

	//只能输入手机号(11位)
	//expreg := "^[1][3-9]\\d{9}$"
	//reg, _ := regexp.Compile(expreg)
	//str := "15067116661"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入手机号(11位)

	//aa,bb := Check_tel(str)
	//fmt.Println(bb)
	//fmt.Println(aa)

	//验证邮箱格式是否正确
	//expreg := "^\\w{1,30}@\\w{1,20}(\\.[a-zA-Z]{1,10}){1,2}$"
	//reg, _ := regexp.Compile(expreg)
	//str := "buzhibujuewoaini@fangchong.com.cn"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//验证邮箱格式是否正确

	//aa,bb := Check_email(str)
	//fmt.Println(bb)
	//fmt.Println(aa)

	//只能输入汉字或字母或数字(1至16个)
	//expreg := "^[\u4e00-\u9fa5\\w]{1,16}$"
	//reg, _ := regexp.Compile(expreg)
	//str := "天市茗大酒店棋牌室1店一二三四五"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入汉字或字母或数字(1至16个)

	//aa,bb := Check_name(str)
	//fmt.Println(bb)
	//fmt.Println(aa)

	//只能输入18位身份证号码，最强验证，年月日严格校验
	//expreg := "^[1-9]\\d{5}(([1][9]\\d{2})|([2][0]\\d{2}))([0][1-9]|[1][012])([0][1-9]|[12][0-9]|[3][01])(\\d{4}|\\d{3}[X])$"
	//reg, _ := regexp.Compile(expreg)
	//str := "11012120551201361X"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入18位身份证号码，最强验证，年月日严格校验


	//aa,bb := Check_IDCard(str)
	//fmt.Println(bb)
	//fmt.Println(aa)


	//只能输入座机电话号码
	//expreg := "^([0]\\d{2,3}-\\d{6,8}(-\\d{1,6})?)$"
	//reg, _ := regexp.Compile(expreg)
	//str := "0571-131234562-123"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入座机电话号码


	//aa,bb := Check_telephone(str)
	//fmt.Println(bb)
	//fmt.Println(aa)


	//只能输入大于0并且小于10000000的正整数或小数(精确0至2位)，刘阳推荐
	//为避免数据库以分为单位存储，所以只能限制7位数。
	//expreg := "^((0\\.(0[1-9]|[1-9]\\d?))|([1-9]\\d{0,6}(\\.\\d{1,2})?))$";
	//reg, _ := regexp.Compile(expreg)
	//str := "1.00"
	//istrue := reg.MatchString(str)
	//istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入大于0并且小于10000000的正整数或小数(精确0至2位)，刘阳推荐
	//
	//
	//aa,bb := Check_money(str)
	//fmt.Println(bb)
	//fmt.Println(aa)


	//只能输入大于或等于1并且小于1000000000的正整数
	//expreg := "^[1-9]\\d{0,9}$";
	//reg, _ := regexp.Compile(expreg)
	//str := "112345678"
	//istrue := reg.MatchString(str)
	////istrue := reg.Match([]byte("123刘阳"))
	//fmt.Println("字符串长度为：",len(str),"匹配结果：",istrue)
	//只能输入大于或等于1并且小于10000的正整数


	//aa,bb := Check_score(str)
	//fmt.Println(bb)
	//fmt.Println(aa)


	//解析url网址为协议、域、端口、从域(路径+文件)、参数、锚
	//expreg := "(\\w+)[:]\\/\\/([^:/]+)([:]\\d*)?([^:?]*)?([^#/]*)?(#.*)?";
	//reg, _ := regexp.Compile(expreg)
	//str := "http://aDmin.fc114.COM.cn:12345/site"
	//arr := reg.FindString(str)
	//fmt.Println(arr)
	//fmt.Println(arr[3:8])
	//解析url网址为协议、域、端口、从域(路径+文件)、参数、锚

	///\w+.*/

	//expreg := "[a-z]+";
	//expreg := "(\\w+)[:]\\/\\/([^:/]+)([:]\\d*)?([^:?]*)?([^#/]*)?(#.*)?";
	//reg, _ := regexp.Compile(expreg)
	//str := "http://aDmin.fc114.COM.cn:12345/site"
	//str := "http://liuyang 15067116661"
	str := "http://liuyang15.067116-661.af.com.com.cn:8080/nihao/asdf-asdflasdf.asdf#asdf?asd=890asdf&add=aa"
	//arr := reg.FindAllString(str,1)
	//fmt.Println(arr)

	istrue := Check_url(str)
	fmt.Println("是否是网址：",istrue)

	uid := "12322"
	isuid := Check_uid(uid)
	fmt.Println("是否是合法的uid：",isuid)

	//检测字符串中是否有数字
	//str := "abc23def01ghi21jk19mno86pqrst11uvw13xyz3619"
	//regStr := "\\d+"
	//result,err := regexp.MatchString(regStr,str)
	//fmt.Println(result,err)//true <nil>
	//检测字符串中是否有数字

	//匹配字符串中的数字
	//str := "abc23def01ghi21jk19mno86pqrst11uvw13xyz3619"
	//regStr := "\\d+"
	//result,err := regexp.Compile(regStr)//解析并返回一个正则表达式
	//byteStr := result.Find([]byte(str))//返回[]byte
	//fmt.Println(byteStr,err)//[50 51] <nil>(ASCII码值：50=2，51=3)
	//
	//stringStr := result.FindString(str)//查找1次匹配，返回string
	//fmt.Println(stringStr,err)//23 <nil>
	//
	//arrStr := result.FindAllString(str,-1)//查找所有匹配，-1不限
	//fmt.Println(arrStr)//[23 01 21 19 86 11 13 3619]
	//匹配字符串中的数字

	//匹配汉字
	//str := "abc23def01ghi21j刘阳k19是mn好o8人6pqrst11uvw13xyz3619"
	//regStr := "[\u4e00-\u9fa5]"//匹配汉字
	//result,err := regexp.Compile(regStr)
	//byteStr := result.Find([]byte(str))//返回[]byte
	//fmt.Println(byteStr,err,"byte转string后：",string(byteStr))//[229 136 152] <nil> byte转string后： 刘
	//
	//stringStr := result.FindString(str)//返回string
	//fmt.Println(stringStr,err)//刘 <nil>
	//
	//arrStr := result.FindAllString(str,-1)//-1不限
	//fmt.Println(arrStr)//[刘 阳 是 好 人]
	//匹配汉字


}

//正则个性化检查，可以是数字开头，可以使用3-20位数字、字母、下划线。
func Check_custom(str string) bool {
	expreg := "^(\\w{3,20})$"
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

//必须以字母开头(参考：微信号和陌陌个性帐号)，可以使用5-20位数字、字母、下划线。
//正则检查用户名只能输入字母或数字或下划线(5至20位)，以字母开头。
func Check_username(str string) bool {
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

//正则检测密码，6-20位。
func Check_password(str string) bool {
	expreg := "^[!@#$%^&*\\w]{6,20}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//正则检测手机号
func Check_mobile(str string) bool {
	expreg := "^[1][3-9]\\d{9}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//正则检测邮箱
func Check_email(str string) bool {
	expreg := "^\\w{1,30}@\\w{1,20}(\\.[a-zA-Z]{1,10}){1,2}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//只能输入汉字、字母、数字、下划线(1至16个)
func Check_title(str string) bool {
	expreg := "^[\u4e00-\u9fa5\\w]{1,16}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//只能输入18位身份证号码，最强验证，年月日严格校验
func Check_IDCard(str string) bool {
	expreg := "^[1-9]\\d{5}(([1][9]\\d{2})|([2][0]\\d{2}))([0][1-9]|[1][012])([0][1-9]|[12][0-9]|[3][01])(\\d{4}|\\d{3}[X])$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//只能输入座机电话号码
func Check_telephone(str string) bool {
	expreg := "^([0]\\d{2,3}-\\d{6,8}(-\\d{1,6})?)$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检测网址是否合法
func Check_url(str string) bool {
	expreg := "^(\\w+)[:]\\/\\/([a-zA-Z0-9][a-zA-Z0-9-]+)[.]([a-zA-Z0-9][a-zA-Z0-9-]+)+([:]\\d*)?([^:?]*)?([^#/]*)?(#.*)?";
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

//检查金额，只能输入大于0并且小于1000000000(9个0)的正整数或小数(精确0至2位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_money(str string) bool {
	expreg := "^((0\\.(0[1-9]|[1-9]\\d?))|([1-9]\\d{0,8}(\\.\\d{1,2})?))$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查金额，只能输入大于0并且小于1000000000(9个0)的正整数或小数(精确0至2位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_s_money(str string) bool {
	expreg := "^(\\+|\\-)?((0\\.(0[1-9]|[1-9]\\d?))|([1-9]\\d{0,8}(\\.\\d{1,2})?))$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查小数，可输入0.00，或输入小于1000000000(9个0)的正整数或小数(精确0至2位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_decimal(str string) bool {
	expreg := "^((0|[1-9]\\d{0,8})(\\.\\d{1,2})?)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}


//有符号(正负)检查小数，可输入0.00，或输入小于1000000000(9个0)的正整数或小数(精确0至2位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_s_decimal(str string) bool {
	expreg := "^((\\+|\\-)?(0|[1-9]\\d{0,8})(\\.\\d{1,2})?)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}



//检查金额，只能输入大于0并且小于1000000000(9个0)的正整数或小数(精确0至6位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_money_6(str string) bool {
	expreg := "^((0\\.(\\d{0,1}[1-9]\\d{0,4}|\\d{0,2}[1-9]\\d{0,3}|\\d{0,3}[1-9]\\d{0,2}|\\d{0,4}[1-9]\\d{0,1}|\\d{0,5}[1-9]|[1-9]\\d{0,5}))|([1-9]\\d{0,8}(\\.\\d{1,6})?))$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查金额，只能输入大于0并且小于1000000000(9个0)的正整数或小数(精确0至6位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_s_money_6(str string) bool {
	expreg := "^(\\+|\\-)?((0\\.(\\d{0,1}[1-9]\\d{0,4}|\\d{0,2}[1-9]\\d{0,3}|\\d{0,3}[1-9]\\d{0,2}|\\d{0,4}[1-9]\\d{0,1}|\\d{0,5}[1-9]|[1-9]\\d{0,5}))|([1-9]\\d{0,8}(\\.\\d{1,6})?))$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查小数，可输入0.000000，或输入小于1000000000(9个0)的正整数或小数(精确0至6位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_decimal_6(str string) bool {
	expreg := "^((0|[1-9]\\d{0,8})(\\.\\d{1,6})?)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查小数，可输入0.000000，或输入小于1000000000(9个0)的正整数或小数(精确0至6位)，刘阳推荐
//注意数据库存储时的精度，也不能用10位，不然会与时间戳相混
func Check_s_decimal_6(str string) bool {
	expreg := "^((\\+|\\-)?(0|[1-9]\\d{0,8})(\\.\\d{1,6})?)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查积分，可输入0，或小于10000000000000000(16个0)的正整数
func Check_s_0_score(str string) bool {
	expreg := "^(\\+|\\-)?(0|[1-9]\\d{0,15})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查积分，可输入0，或小于10000000000000000(16个0)的正整数
func Check_0_score(str string) bool {
	expreg := "^(0|[1-9]\\d{0,15})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查积分，只能输入大于或等于1并且小于10000000000000000(16个0)的正整数
func Check_1_score(str string) bool {
	expreg := "^([1-9]\\d{0,15})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查数字，可输入0，或小于100000(5个0)的正整数
func Check_s_0_100000(str string) bool {
	expreg := "^(\\+|\\-)?(0|[1-9]\\d{0,4})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}


//检查数字，可输入0，或小于100000(5个0)的正整数
func Check_0_100000(str string) bool {
	expreg := "^(0|[1-9]\\d{0,4})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查数字，只能输入大于或等于1并且小于100000(5个0)的正整数
func Check_1_100000(str string) bool {
	expreg := "^([1-9]\\d{0,4})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检查数字，可输入0，或小于1000000000(9个0)的正整数
func Check_s_0_1000000000(str string) bool {
	expreg := "^(\\+|\\-)?(0|[1-9]\\d{0,8})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查数字，可输入0，或小于1000000000(9个0)的正整数
func Check_0_1000000000(str string) bool {
	expreg := "^(0|[1-9]\\d{0,8})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查数字，只能输入大于或等于1并且小于1000000000(9个0)的正整数
func Check_1_1000000000(str string) bool {
	expreg := "^([1-9]\\d{0,8})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查数字，可输入0，或小于uint64最大值的正整数
func Check_0_20(str string) bool {
	expreg := "^(0|[1-9]\\d{0,19})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//只能输入大于或等于1并且小于uint64最大值的正整数
func Check_1_20(str string) bool {
	expreg := "^([1-9]\\d{0,19})$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//正则检测字符串中的内容是否为纯数字，不限制个数
func Check_number(str string) bool {
	expreg := "^(\\d+)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//有符号(正负)检测字符串中的内容是否为正数或负数
func Check_s_number(str string) bool {
	expreg := "^((\\+|\\-)?\\d+)$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}


//正则检测用户uid格式是否合法，以大于0值开头的数字且5到19位为合法。
func Check_uid(str string) bool {
	expreg := "^[1-9]\\d{4,18}$"
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}

//检查6位数字
func Check_6(str string) bool {
	expreg := "^[0-9]{6}$";
	reg, regErr := regexp.Compile(expreg)
	if regErr != nil {
		return false
	}
	istrue := reg.MatchString(str)

	return istrue
}


//正则检测md5格式是否合法，判定长度为固定32位的数字和字母
func Check_md5(str string) bool {
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

//正则检测token格式是否合法，判定为30至100之间的数字和字母
func Check_token(str string) bool {
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

//正则检测jwt格式是否合法，判定为50至300之间的任意字符(不包括',")
func Check_jwt(str string) bool {
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

//正则检测设备id格式是否合法，判定为6至200之间的任意字符(不包括',")
func Check_devid(str string) bool {
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
