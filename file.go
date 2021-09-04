package liuyang

import (
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strconv"
)

//将数值转为文件类型数值，如：KB、MB、GB、TB、HB为单位的数值字符串。
func File_size(filesize int64) string {
	var KB int64 = 1024		//1KB
	MB := KB * 1024			//1MB
	GB := MB * 1024			//1GB
	TB := GB * 1024			//1TB
	PB := TB * 1024			//1PB
	EB := PB * 1024			//1EB
	/*
		1024							4		1KB
		1048576							7		1MB
		1073741824						10		1GB
		1099511627776					13		1TB
		1125899906842624				16		1PB
		1152921504606846976				19		1EB
	*/
	var temp float64 = 0
	//不足1KB时，以B结尾
	if filesize < KB {//这里必须要判断小于，不能判断小于等于，不然会出现1024B的情况。

		return strconv.FormatInt(filesize,10) + "B"

	} else if filesize < MB {//如果小于MB，也就是大于KB，而小于MB时，就除以KB，并保留两位小数。

		temp = float64(filesize) / float64(KB)
		return Float64ToString(temp,2) + "KB"

	} else if filesize < GB {

		temp = float64(filesize) / float64(MB)
		return Float64ToString(temp,2) + "MB"

	} else if filesize < TB {

		temp = float64(filesize) / float64(GB)
		return Float64ToString(temp,2) + "GB"

	} else if filesize < PB {

		temp = float64(filesize) / float64(TB)
		return Float64ToString(temp,2) + "TB"

	} else if filesize < EB {

		temp = float64(filesize) / float64(PB)
		return Float64ToString(temp,2) + "PB"

	} else if filesize < math.MaxInt64 {

		temp = float64(filesize) / float64(EB)
		return Float64ToString(temp,2) + "EB"

	}

	return "数值" + strconv.FormatInt(filesize,10) + "超过int64的最大长度，无法解析。"
}


//文件操作===========================================
//将字符串写入文件中，追加（文件内容尾部追加），返回写入的字节数和可能遇到的任何错误
func File_put_contents(filename string,data string) (int, error) {
	//判断文件是否存在
	if File_exits(filename) == false {//文件不存在
		//不存在，创建新的文件，创建新文件的前提条件是，其上的目录必须存在才行，所以要保证其目录已存在。
		//首先要创建目录，只有事先创建好目录，才能在此目录下创建文件，而创建目录的前提，是要先获取filename路径中的目录，所以：
		File_mkdir(filename)
	}
	//已存在时调用文件写操作方法
	//文件尾部追加内容的模式打开文件
	fopen,ferr := os.OpenFile(filename,os.O_APPEND|os.O_WRONLY|os.O_CREATE,os.ModePerm)//0666或os.ModeAppend也可以
	CheckErr(ferr)
	defer fopen.Close()//关闭文件
	return fwrite(fopen, data)
}

//文件写操作
func fwrite(fopen *os.File, content string) (int, error) {

	//调用系统内置方法，将字符串写入文件中
	wbyte,werr := fopen.WriteString(content)

	return wbyte, werr
}

//将字符串写入文件中，重写（清空文件内容），返回写入的字节数和可能遇到的任何错误
func File_put_contents_rewrite(filename string,data string) (int, error) {
	//判断文件是否存在
	if File_exits(filename) == false {//文件不存在
		//不存在，创建新的文件，创建新文件的前提条件是，其上的目录必须存在才行，所以要保证其目录已存在。
		//首先要创建目录，只有事先创建好目录，才能在此目录下创建文件，而创建目录的前提，是要先获取filename路径中的目录，所以：
		File_mkdir(filename)
		//第三步，创建文件
		FCcreate,FCerr := os.Create(filename)
		CheckErr(FCerr)//检查错误
		defer FCcreate.Close()//关闭句柄
		//调用文件写操作方法
		return fwrite(FCcreate, data)
	}
	//存在时调用文件写操作方法
	//以只写的方式打开文件，如果可能，打开时清空文件，0666代表读和写的权限位。
	fopen,ferr := os.OpenFile(filename,os.O_WRONLY|os.O_TRUNC,0666)
	CheckErr(ferr)//检查错误
	defer fopen.Close()//关闭文件

	return fwrite(fopen, data)
}


//二次封装，读取文件内容，将读到的[]byte数据用高级写法转成string，并和bool值一起返回
func File_get_contents(filename string) (string, bool) {
	
	bb,istrue := File_get_data(filename)
	if istrue == false {
		return "", false
	}
	//将[]byte高级转法为string
	return BytesToString(bb), true
}

//读取文件内容，返回[]byte数据和bool值。
func File_get_data(filename string) ([]byte, bool) {
	//判断文件是否不存在
	if File_exits(filename) == false {//文件不存在
		//文件不存在
		return nil, false
	}
	//存在时
	fopen,ferr := os.Open(filename)//打开文件
	CheckErr(ferr)//检查错误
	defer fopen.Close()//关闭文件
	//获取文件信息
	Sinfo,_ := fopen.Stat()//这里可以直接使用open的句柄调用Stat方法。而不用os.Stat方法。
	//调用文件读操作方法，传参文件句柄和文件字节大小。
	return fread(fopen, Sinfo.Size()), true
}

//读取文件操作，返回[]byte
func fread(fopen *os.File, length int64) []byte {

	//创建byte类型的数组
	byteArr := make([]byte,length)

	//调用系统内置方法，读取文件中的内容
	rbyte,rerr := fopen.Read(byteArr)//往这个[]byte数组中读，也就是将读到的数据存储到[]byte数组中。
	//liuyang.txt文件中的内容为：abc02刘阳是好人，所以下面的打印是20个字节数，也就是每个汉字占3个字节数的长度
	// Read方法可忽略，只有ReadAt方法在读取文件时，如果设置要读取的长度byte大于文件内容的长度时，读完后会出现EOF，所以当EOF的时候不做处理。
	CheckErrEOF(rerr)//当读取文件出错时
	//将读取到的内容长度用byte数组进行读取并转成字符串
	return byteArr[0:rbyte]//普通转法
	//return string(byteArr[0:rbyte])//普通转法
	//return BytesToString(byteArr[0:rbyte])//高级转法
}
//二次封装，读取文件内容，参数1，要读取的文件名称及路径，参数2，要读取的开始位置，参数3，要读取的长度，将读取的[]byte用高级写法转成string后和bool一起返回
func File_get_contents_index(filename string, index int64, length int64) (string, bool) {

	bb,istrue := File_get_data_index(filename, index, length)
	if istrue == false {
		return "", false
	}

	//将[]byte高级转法为string
	return BytesToString(bb), true
}

//读取文件内容，参数1，要读取的文件名称，参数2，要读取的开始位置，参数3，要读取的长度，返回读取到的[]byte数据和bool值
func File_get_data_index(filename string, index int64, length int64) ([]byte, bool) {
	//判断文件是否不存在
	if File_exits(filename) == false {//文件不存在
		//文件不存在
		return nil, false
	}
	//存在
	fopen,ferr := os.Open(filename)//打开文件
	CheckErr(ferr)//检查错误
	defer fopen.Close()//关闭文件
	//获取文件信息
	Sinfo,_ := fopen.Stat()
	//得到文件内容的字节大小
	//当传过来要读取文件的长度大于文件内容的长度时，将文件内容的长度赋值给要读取的长度
	if length > Sinfo.Size() {
		length = Sinfo.Size()
	}
	//调用按照索引位置和长度进行文件读操作方法
	return fread_index(fopen,index,length), true
}

//读取文件内容操作，参数1，文件名柄指针，参数2，要读取的开始位置，参数3，要读取的长度，返回[]byte数据
func fread_index(fopen *os.File,index int64, length int64) []byte {

	//创建byte类型的数组
	byteArr := make([]byte,length)

	//调用系统内置方法，读取文件中的内容
	rbyte,rerr := fopen.ReadAt(byteArr,index)//往这个[]byte数组中读，也就是将读到的数据存储到[]byte数组中。
	//liuyang.txt文件中的内容为：abc02刘阳是好人，所以下面的打印是20个字节数，也就是每个汉字占3个字节数的长度
	//ReadAt方法在读取文件时，如果设置要读取的长度byte大于文件内容的长度时，读完后会出现EOF，所以当EOF的时候不做处理。
	CheckErrEOF(rerr)//当读取文件出错时
	//将读取到的内容长度用byte数组进行读取并转成字符串
	//return string(byteArr[0:rbyte])//为了将byte转string更高效的使用，所以这里返回[]byte而不返回string
	return byteArr[0:rbyte]
}

//检查文件或目录是否存在，如果存在返回true，否存返回false
func File_exits(path string) bool {
	_,ferr := os.Stat(path)//只有不存在时，ferr才有内容，否则是nil
	//这里的ferr在不等于nil时，也有可能是"aa/bb"这样的字符串，而没有实际目录文件。所以不能以是否为nil作为判断依据。
	//if ferr != nil {
	//	return true,nil
	//}
	//IsNotExist函数判断的是文件不存在。存在返回的是false，不存才返回true。
	if os.IsNotExist(ferr) == false {
		return true//存在
	}
	return false//不存在
}

//得到文件信息
func File_info(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

//获取路径中的目录并创建
func File_mkdir(filename string) error {
	//第一步，获取目录
	//文件路径为liuyang/sunlulu/nihao/aa.jpg时会得到：liuyang/sunlulu/nihao
	dirname := path.Dir(filename)
	//第二步，创建目录。
	//MkdirAll函数内部已经判断参数是否是目录，如果是就不用创建。所以这里就不用再判断了。
	mkerr := os.MkdirAll(dirname, os.ModePerm)//创建多级目录

	return mkerr
}

//文件操作===========================================

//断点续传，参数1：源文件、参数2：目标文件
func File_xu_chuan(src_file string, desc_file string) error {

	//测试复制断点续传

	//源文件名称
	//src_file := "sunlulu/meinv.png"
	//目标文件名称
	//desc_file := "meinv/newfile.jpg"
	//临时文件名称
	temp_filename := desc_file + "temp.txt"

	//fmt.Println(src_filename)//sunlulu/meinv.png
	//fmt.Println(desc_filename)//meinv/newfile.jpg
	//fmt.Println(temp_filename)//meinv/newfile.jpgtemp.txt

	//创建源文件句柄，只读模式
	srcFile,srcErr := os.Open(src_file)
	CheckErr(srcErr)//检查错误
	defer srcFile.Close()//关闭已打开的源文件句柄

	//创建目标文件句柄，只写模式，没有则创建
	File_mkdir(desc_file)//首先要创建目录
	descFile,descErr := os.OpenFile(desc_file,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	CheckErr(descErr)//检查错误
	defer descFile.Close()//关闭已打开的目标文件句柄

	//创建临时文件句柄，读写模式，没有则创建
	tempFile,tempErr := os.OpenFile(temp_filename,os.O_RDWR|os.O_CREATE,os.ModePerm)
	CheckErr(tempErr)//检查错误


	//声明读取源文件数据后所存放的容器，并设置每次读取的数据量
	read_srcData := make([]byte,1024)
	//声明读取临时文件数据后所存放的容器，并设置每次读取的数据量
	read_tempData := make([]byte,19)//19位数是int64类型数据的最大实际位数

	//声明已写入累加的数据总量
	var total int64 = 0

	//开始循环复制文件
	for {
		//第1步，设置临时文件句柄的偏移量（0，文件开头）。
		tempFile.Seek(0, io.SeekStart)

		//第2步，读取临时文件中的数据。
		temp_int,temp_err := tempFile.Read(read_tempData)
		CheckErrEOF(temp_err)

		//第3步，将读到临时byte容器中的数据转成string类型，容器数据取：0到已读字节数。如果不指定范围，会转换为空。
		str_read_tempData := string(read_tempData[0:temp_int])

		//第4步，将临时文件中的内容，也就是刚刚转换的字符串，转成int64类型。该值为：总计已读数。
		total,_ = strconv.ParseInt(str_read_tempData,10,64)

		//第5步，设置源文件句柄偏移量(总计已读数，文件开头)，准备读取源文件数据。
		srcFile.Seek(total,io.SeekStart)

		//第6步，读取源文件数据，得到返回已读字节数。
		read_int,read_err := srcFile.Read(read_srcData)
		CheckErrEOF(read_err)//检查错误，并排除EOF
		//fmt.Println(read_int, read_err)
		//打印传输过程
		//1024 <nil>
		//1024 <nil>
		//1024 <nil>
		//646 <nil>
		//0 EOF

		//第7步，判断源文件已读字节数为0时并且返回错误类型为EOF时，证明源文件数据已读到末尾，所以break跳出循环。
		//在跳出循环之前有两步操作：1，关闭临时文件句柄；2，删除临时文件。
		if read_int == 0 && read_err == io.EOF {
			//关闭临时文件打开的句柄
			tempFile.Close()//关闭已打开的临时文件句柄

			//删除临时文件
			os.Remove(temp_filename)//删除临时文件

			break//停止特环
		}

		//第8步，设置目标文件句柄偏移量(总计已读数，文件开头)，准备写入已读取的数据，此步骤针对：O_WRONLY，如果openfile打开模式为append时，当前步骤可省略。
		descFile.Seek(total,io.SeekStart)

		//第9步，向目标文件中写入数据。写入的数据为：读取源文件存储的容器，长度为0到读取源文件后返回的字节数。
		write_int,write_err := descFile.Write(read_srcData[0:read_int])
		CheckErr(write_err)

		//第10步，将写入的数据量累加。临时文件中存储的数据加上刚刚已写入到目标文件中的数据量。等于已复制的总量值。
		total += int64(write_int)

		//第11步，将已写入的实际数据量累加值转换为string类型。
		str_total := strconv.FormatInt(total,10)

		//第12步，设置临时文件句柄偏移量（0，文件开头），准备写入已写入目标文件后所返回的实际写入量。
		tempFile.Seek(0,io.SeekStart)

		//第13步，将string类型的已写入数据量写入临时文件中。
		_,temp_write_err := tempFile.WriteString(str_total)
		CheckErr(temp_write_err)//检查写入临时文件后的错误返回

		//fmt.Println("当前已复制数据为：",total)
		//第14步，可模拟断电。只需判断：总计已读数大于等于某一值时，panic即可。
		//if total >= 10000 {
		//	panic("测试断电")
		//}
	}
	//开始循环复制文件

	//测试复制断点续传

	return nil
}



//统计目录中的文件夹数量与文件数量，参数1为：目录地址，参数2为：层次计数
//返回：文件夹数量累加递规返回的文件夹数量，文件数量累加递规返回的文件数量
func File_folder_total(dir string, level int) (int, int){
	fileInfo_arr,fileInfo_err := ioutil.ReadDir(dir)
	CheckErr(fileInfo_err)
	fcount := len(fileInfo_arr)//当前目录中文件夹与文件的数量

	level_str := "|--"

	//循环层次，有几层，就循环几次，如有6层，那就打印6个空格。此处只是把6个空格连接上。以便在下面打印时使用。
	for k := 0; k < level; k++ {
		level_str = "|   " + level_str
	}

	//此处如果不用的话，那么递规调用方法时就要传参level + 1这种形式，其实和level++功能一样。
	//level++ //层次变量自增，想要打印有层次感的文件名地址，就必须要有该变量的记录。

	child_dir := ""		//子文件夹变量

	var folder_num int = 0 //文件夹数量
	var file_num int = 0 //文件数量

	//统计子文件夹中的文件夹数量和文件数量
	n_folder := 0		//文件夹数量
	n_file := 0			//文件数量

	for i := 0; i < fcount; i++ {
		//fmt.Println("共：",fcount,"个文件，","文件名为：",fileInfo_arr[i].Name(),"，是否目录：",fileInfo_arr[i].IsDir())
		if fileInfo_arr[i].IsDir() == true {//如果是文件夹时，则继续遍历
			child_dir = dir + "/" + fileInfo_arr[i].Name()
			folder_num++	//文件夹数量累加
			//递规调用，参数为目录地址和层次，层次如果不加1的话，那上面就需要level++了。
			n_folder,n_file = File_folder_total(child_dir,level + 1)//返回文件夹数量和文件数量
			folder_num += n_folder	//文件夹数量累加递规返回的文件夹数量
			file_num += n_file		//文件数量累加递规返回的文件数量
			//fmt.Println("文件夹数量为：",folder_num,",第2值：",n_folder,"，文件数量",n_file,"，文件夹名称为：",child_dir)
		} else {
			file_num++				//文件数量累加
			//fmt.Println("文件为：",fileInfo_arr[i].Name())
		}

		//tempdir := dir + "/" + fileInfo_arr[i].Name()
		//fmt.Println("当前级别：",level,"，文件为：",level_str,tempdir)

	}

	return folder_num, file_num
}
//统计目录中的文件夹数量与文件数量
