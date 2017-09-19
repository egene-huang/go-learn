package main

import "fmt"

func main() {
	//
//	rc()
	rc1()
}

func rc() {
	var a, b = fn()
	fmt.Println(a,b)
}

func fn() (int,int) {
	//函数定义: func func_name(arguments args)[return_type] {}
	//函数返回多个值的时候,需要在函数返回值类型处使用()将多个返回值类型包裹, 并使用,分开
	//函数可以同时返回多个值
	return 3,4
}

func rc1() {

	var _, x = fn() //_代表丢弃对应的返回值 不能读取改值  这里3被丢弃
	fmt.Println(x)
}
