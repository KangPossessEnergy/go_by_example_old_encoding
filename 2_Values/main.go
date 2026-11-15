package main

import "fmt"

// Go 拥有多种值类型，包括字符串、整型、浮点型、布尔型等。
func main() {

	//1 字符串，可以通过 + 连接
	fmt.Println("go" + "lang") //golang

	//2 整数和浮点数
	fmt.Println("1+1=", 1+1)          //1+1= 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) //7.0/3.0 = 2.3333333333333335

	//3 布尔型

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         //false

}
