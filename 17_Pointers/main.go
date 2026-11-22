package main

import "fmt"

/*
指针
go语言中一种特殊的数据类型，它存储的是另一个变量的内存地址，而不是变量本身的值
简单理解：
普通变量：存储数据本身
指针变量：存储数据所在的内存地址
*/
func main() {
	case1()

	//Go 支持 指针， 允许在程序中通过 引用传递 来传递值和数据结构。

	i := 1
	fmt.Println("initial:", i) //initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) //zeroval: 1

	//通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i) //zeroptr: 0

	//指针也是可以被打印的
	fmt.Println("pointer:", &i) //pointer: 0x691f0dae040

	// zeroval 在 main 函数中不能改变 i 的值， 但是 zeroptr 可以，因为它有这个变量的内存地址的引用。
}

// 我们将通过两个函数：zeroval 和 zeroptr 来比较 指针 和 值。

// zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival。
func zeroval(ival int) {
	ival = 0
}

// zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。
// 紧接着，函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。 对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func case1() {

	a := 100

	p := &a
	fmt.Println("a的值:", a)    //a的值: 100
	fmt.Println("a的地址:", &a)  //a的地址: 0x3de5df4460d8
	fmt.Println("p存储的地址:", p) //p存储的地址: 0x3de5df4460d8
	fmt.Println("p指向的值:", *p) // p指向的值: 100

	/*
		三个重要符号
		&：获取变量的内存地址.   eg: &a
		*: 根据地址获取变量的值（解引用）.  eg: *p
		*int：表示一个指向int类型的指针. eg：var p *int

	*/

}
