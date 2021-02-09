package liuyang

import "io"

//用于检测错误，如出现错误会打印该错误并中止程序运行
func CheckErr(err error) {
	if err != nil {
		panic(err)
		//log.Fatal(err)//2019/08/20 13:05:41 strconv.Atoi: parsing "2c": invalid syntax
	}
}

//用于检测错误，并且排除EOF，如出现错误并且不是EOF时会打印该错误并中止程序运行
func CheckErrEOF(err error) {
	if err != nil && err != io.EOF {
		panic(err)
		//log.Fatal(err)//2019/08/20 13:05:41 strconv.Atoi: parsing "2c": invalid syntax
	}
}

