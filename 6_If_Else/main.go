package main

import "fmt"

// if/else 分支
func main() {

	//if 和 else 分支结构在 Go 中非常直接。

	//1 基本的例子

	//7对2取余，为1
	if 7%2 == 0 {
		fmt.Println("7 is even") // even:偶数
	} else {
		fmt.Println("7 is odd") //odd：奇数
	}

	//2 你可以不要 else 只用 if 语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4 ") //8 is divisible by：8能被4整除
	}

	//3 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative") //is negative：是负面的
	} else if num < 10 {
		fmt.Println(num, "has 1 digit") //has 1 digit：有一位数字
	} else {
		fmt.Println(num, "has multiple digits") //has multiple digits：具有多个数字
	}

	/*
		注意：
		1 在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。
		2 Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。


	*/
}
