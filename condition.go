package main

import "fmt"

func main() {
	//条件判断 if后面紧跟的只要是布尔值(bool)就行, 小括号可以没有
	cond(19)
	cond1(20)
	cond2()
}

func cond(xx int) {
//条件判断 
	if xx > 20 {
		fmt.Println("大于20")
	}else {
		fmt.Println("不大于20")	
	}
}

func cond1(xx int) {
//这里可以有小括号
	if(xx > 20) {
		fmt.Println("大于20")
	}else {
		fmt.Println("小于20")
	}
}

func cond2() {
	var xx bool //布尔值 默认为false
	if xx {
		fmt.Println("通过.")
	}else {
		fmt.Println("不通过")
	}
}
