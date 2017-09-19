package main

import "fmt"

func main() {
	str := "Hello GO Language."
	fmt.Println(str)
	fn()
	set()
	eq()
}

func init() {
	fmt.Println("初始化函数.")	
}

func fn() {
	var str string = "这是另一种变量申明方式"
//局部申明变量必须使用 编译报错
	fmt.Println("Hello. ", str)
}

func set() {
	//多个变量初始化, 自动对应赋值, 和python类似
	a, b, c := 1,"Hello", 0.99999
	fmt.Println(a,b,c)
}

func eq() {
	var a, b = "hello", "hello"
	fmt.Println(a == b) //相等
	fmt.Println(&a)//返回变量a存储地址  0xc42000e200
}
