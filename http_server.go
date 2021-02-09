package liuyang

//func HttP_Server_test()  {
//	http.HandleFunc("/json",HTTPResponseReceiveJSON)
//	http.ListenAndServe(":8081",nil)
//}


/*

//接收

//http接收post请求
func HTTPResponseReceive(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parse form error ",err)
	}
	// 初始化请求变量结构
	formData := make(map[string]interface{})
	result,rerr := ioutil.ReadAll(r.Body)
	fmt.Println(rerr)
	fmt.Println(string(result))
	// 调用json包的解析，解析请求body
	json.NewDecoder(r.Body).Decode(&formData)
	for key,value := range formData{
		fmt.Println("key:",key," => value :",value)
	}
	io.WriteString(w,string(result))
}

//http接收post请求，并返回json数据
func HTTPResponseReceiveJSON(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal("parse form error ",err)
	}

	//解析客户端提交过来的json数据
	pp1 := &Person{}
	//str,_ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(str,pp)
	//fmt.Println(pp)//&{刘阳 25 男 [{25.56 3} {6.8 2}]}
	//zhi := pp.Tels[1].Money
	//fmt.Println(zhi)//6.8
	//fmt.Println("========================")
	json.NewDecoder(r.Body).Decode(pp1)
	fmt.Println(pp1)//&{刘阳 25 男 [{25.56 3} {6.8 2}]}
	fmt.Println(pp1.Tels[0].Money)//25.56
	// 初始化请求变量结构
	//formData := make(map[string]interface{})
	//// 调用json包的解析，解析请求body
	//json.NewDecoder(r.Body).Decode(&formData)
	//fmt.Println(formData["tels"])
	//for key,value := range formData{
	//	log.Println("key:",key," => value :",value)
	//}


	//返回json字符串给客户端
	//pp := &Person{}
	//formData := make(map[string]interface{})
	w.Header().Set("Content-Type","application/json")
	//json.NewEncoder(w).Encode(formData)//返回自定义的map数据给客户端
	//json.NewEncoder(w).Encode(pp)//返回结构体数据也就是json字符串给客户端
	//如果是一维数据，可以用自己定义的map数据解析客户端传来的json。因为碰到二维数据时会出现问题。
	//如果是一维包含二维数据或多维数据时，就要定义和客户端一样的结构体进行解析，这样才能保证其数据和客户端一样。
	//所以推荐使用第2种。

	pp := &Person{
		User: "孙璐璐",
		Age:  22,
		Sex:  "女",
		Tels: []Tel{
			{
				Money:16.58,
				Num:3,
			},
			{
				Money:8.36,
				Num:2,
			},
		},
	}
	//第一种方式：
	//jsonByte, _ := json.Marshal(pp)
	//io.WriteString(w,string(jsonByte))

	//第二种方式：
	json.NewEncoder(w).Encode(pp)
}
*/
