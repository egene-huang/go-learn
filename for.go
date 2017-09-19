package main

import "fmt"

func main() {
	//
//	fo(10)
//	fo1(5, 10)
//	fo2()
	gt()
}

func fo(n int) {
	//这里不能有小括号包裹
	for x := 0; x < n; x++ {
		fmt.Println("第", x + 1, "次循环")
	}
}


func fo1(n int, m int) {
	x, y := n, m
	for x < y { //这种for循环 类似 for 内部if 结合break使用
		x ++ 
		fmt.Println("x", x, y)
	}

	for n < m {
		n += 5
		fmt.Println("----")	
	}
}

func fo2() {
	for n := 0; n < 100; n++ {
		if n == 5 {
			continue ;		
		}

		if n == 50 {
			break ;
		}
		
		fmt.Println("打印 ---" , n)
	}
}

func gt() {

	lable: for x := 0; x < 10; x++ {
		if x == 5 {
			//构造死循环 一直打印0-4  和java的 goto实现是有区别的 区别之一在于 局部变量不会被重新初始化
			goto lable		
		}
		fmt.Println("goto[打印",x)
	}
}
