package main

import (
	"fmt"
	"math"
)

// Go 支持字符、字符串、布尔和数值 常量 。

const s string = "constant"

func main() {

	//const 语句可以出现在任何 var 语句可以出现的地方
	fmt.Println(s) //constant

	//常数表达式可以执行任意精度的运算
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d) //6e+11

	//数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
	//int64(d) —— 显式类型转换
	fmt.Println(int64(d)) //600000000000

	//一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。 举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型。
	//math.Sin(n) —— 自动确定类型
	fmt.Println(math.Sin(n)) //-0.28470407323754404
}
