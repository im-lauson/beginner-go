package main

//func user(w http.ResponseWriter,r *http.Request)  {
//	fmt.Fprintf(w,format:"dss %s",r.URL.Path[1:])
//}
/*1--1.Go 基本语法和 Web 框架起步*/
func main() {

	//http.HandleFunc(pattern:'/',user)
	//log.Fatal(http.ListenAndServe(addr:":8080",handler:nil))

	var lan string = "Hello World!"
	print(lan)
	//rune()是一个int32,一个rune四个字节，私下可以看作字符
	print(len(lan))
	//var a = "aa"重复声明变量不会通过语法和编译
	var a = "bb"
	print(a)
	//len,长度
	//const常量的声明
	/*类型总结：
	bool 表示 boolean 类型，用来表达 true 或 false 的。
	int：代表 32 或 64 位整数，具体取决于基础平台。除非需要使用特定大小的整数，否则通常应该使用
	uint： 表示 32 或 64 位无符号整数，具体取决于基础平台。
	float32：32 位浮点数
	float64：64 位浮点数
	complex64：具有 float32 实部和虚部的复数
	complex128：具有 float64 实部和虚部的复数
	byte：uint8 的别名
	rune：int32 的别名
	字符串类型：在 Go 语言中，字符串是字节的集合。如果这个定义没有任何意义，那也没关系。现在，我们可以假设字符串是字符的集合。

	*/
	_, age, sex := Name()
	println(age, sex) //用_忽略返回值，标记为匿名变量
	println(Name())   //方法的调用

}

// Name 方法的声明（多变量支持多返回值）   		返回值类型
func Name() (name string, age int, sex bool) {
	return "Lauson", 12, true //参数命名，返回值不命名

}
