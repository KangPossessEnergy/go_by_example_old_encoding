package main

import "fmt"

// 在 Go 中，变量 需要显式声明，并且在函数调用等情况下， 编译器会检查其类型的正确性。
func main() {
	//1 var 声明 1 个或者多个变量。
	var a = "initial"
	fmt.Println(a)

	//2 一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c) //1 2

	//3 Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d) //true

	//4 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。 例如，int 的零值是 0。
	var e int
	var f string
	fmt.Println("e", e) //0
	fmt.Println("f", f) //

	//5 := 语法是声明并初始化变量的简写， 例如 var g string = "kkdw" 可以简写为右边这样。
	g := "kkdw"
	fmt.Println(g) //kkdw
}
