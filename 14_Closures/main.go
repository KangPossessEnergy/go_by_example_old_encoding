package main

import "fmt"

// Go 支持匿名函数， 并能用其构造 闭包。 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。

// 1 intSeq 函数返回一个在其函数体内定义的匿名函数。返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包。
func intSeq() func() int {
	// i := 0
	var i = 0
	return func() int {
		i++
		return i
	}
}

// 闭包
func main() {
	//我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
	nextInt := intSeq()

	//通过多次调用 nextInt 来看看闭包的效果。
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	//为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下
	newInts := intSeq()
	fmt.Println(newInts()) //1

}
