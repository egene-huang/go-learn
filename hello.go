/* 这是注释 不会执行*/
package main
import "fmt"

//每个go程序都会有一个main函数
func main() {
	var f int = 0 //申明一个int类型变量 f 使用 var申明
//在申明的时候可以不指定变量类型, go自行判断
	var x = 0.0
	def(f) //自定义函数必须在main函数中调用
	fn(x)
	fun(0.2)
	a := 123
	fmt.Println(a)
	fmt.Println("Hello, Arch")
}

//这是注释, 不会执行, init函数会先于main函数执行.
func init() {
	fmt.Println("这是init函数.")
}

//大写字母开头的申明可以被包外部代码访问
//一行代表一个语句结束
func def(xx int) { //申明形参 xx
	fmt.Println(xx) 
	//没有类似java string类型 运算符重载 +
	fmt.Println("in order to.... BlockChain....") 
}

func fn(xx float64) {
	fmt.Println(xx) 
}

func fun(x float64) {
	f := x //不使用var 申明 go自行推导  f不能是已经申明过的变量
//这种形式只能用在函数内
	fmt.Println(f)
}


