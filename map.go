package liuyang

import (
	"strconv"
)

//将struct类型数据转换成map类型数据，转不了，不要再试了。也没必要转，直接将struct转成map岂不更好。

//将map数据转成post提交所需格式的字符串
func MapToString(mm map[string]interface{}) string {
	//样例：name=65.56&sex=25&age=true&
	//将最后面的字符"&"去掉，有两种方式：

	//方式1：
	//str := ""//总接收变量
	//temp := ""//临时接收拼接的变量
	//for key,value := range mm {
	//	zhi := ValueToString(value)
	//	temp = key + "=" + zhi + "&"
	//	str += temp
	//}
	////由于map的key值顺序是不固定的，所以无法判断最后一个key是什么值，所以只能用最后出现的"&"符号来定位。
	//index := strings.LastIndex(str,"&")
	//result := str[0:index]
	//return result



	//方式2：
	str := ""//总接收变量
	temp := ""//临时接收拼接的变量
	i := 0//自增变量
	cnt := len(mm)//map元素个数
	for key,value := range mm {
		zhi := ValueToString(value)
		if i < (cnt - 1) {
			temp = key + "=" + zhi + "&"
		} else {
			temp = key + "=" + zhi
		}
		str += temp
		i++//临时变量自增
	}
	return str
}

//interface值转string
func ValueToString(i interface{}) string {
	//fmt.Println(i)//打印参数值

	str := ""
	//用这种方法判断，就省去了reflect.TypeOf(i)反射的判断，如：
	//obj := reflect.TypeOf(i)
	//if obj.Kind() == reflect.Int {}
	switch idata := i.(type) {
	case string:
		str = idata
	case int:
		str = strconv.Itoa(idata)
	case int8:
		str = strconv.Itoa(int(idata))
	case int16:
		str = strconv.Itoa(int(idata))
	case int32:
		str = strconv.Itoa(int(idata))
	case int64:
		str = strconv.FormatInt(idata,10)
	case uint:
		str = strconv.Itoa(int(idata))
	case uint8:
		str = strconv.Itoa(int(idata))
	case uint16:
		str = strconv.Itoa(int(idata))
	case uint32:
		str = strconv.FormatInt(int64(idata),10)
	case uint64:
		str = strconv.FormatInt(int64(idata),10)
	case float32:
		str = strconv.FormatFloat(float64(idata),'f',-1,32)
	case float64:
		str = strconv.FormatFloat(idata,'f',-1,64)
	case bool:
		str = strconv.FormatBool(idata)
	case []byte:
		str = string(idata)
	default:
		str = "error"//未知类型时，返回error字符串
	}

	return str
}

