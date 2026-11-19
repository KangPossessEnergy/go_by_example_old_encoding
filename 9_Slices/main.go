package main

import (
	"fmt"
	"reflect"
)

// 切片，Slice 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。
func main() {

	//1 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。
	//要创建一个长度不为 0 的空 slice，需要使用内建函数 make。 这里我们创建了一个长度为 3 的 string 类型的 slice（初始值为零值）。

	// var s = make([]string, 3)
	s := make([]string, 3)
	fmt.Println("s", s) //s [  ]

	//2 可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//3 len 返回 slice 的长度
	fmt.Println(len(s)) //3

	//4 除了基本操作外，slice 支持比数组更丰富的操作。
	//  比如 slice 支持内建函数 append， 该函数会返回一个包含了一个或者多个新值的 slice。
	//  注意由于 append 可能返回一个新的 slice，我们需要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) //apd: [a b c d e f]

	//5 slice 还可以 copy。这里我们创建一个空的和 s 有相同长度的 slice——c， 然后将 s 复制给 c。
	c := make([]string, len((s)))
	copy(c, s)
	fmt.Println("copy:", c) //copy: [a b c d e f]

	//6 slice 支持通过 slice[low:high] 语法进行“切片”操作。 例如，下边的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
	l := s[2:5]
	fmt.Println("l", l) //l [c d e]

	//7 这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素。
	l = s[:5]
	fmt.Println("sl2:", l)

	//8 这个 slice 包含从 s[2]（包含 2）之后的元素。
	l = s[2:]
	fmt.Println("sl3:", l)

	//9 我们可以在一行代码中声明并初始化一个 slice 变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) //dcl: [g h i]

	//10 Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
	//make()：创建并初始化切片
	//[][]int：二维切片类型
	//twoD[i][j]：访问二维切片中的元素
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) //创建一个长度为 innerLen 的整数切片，并自动初始化为 0。 twoD[0] = [0]
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d", twoD) //2d [[0] [1 2] [2 3 4]]

	//11 注意，slice 和数组是不同的类型，但它们通过 fmt.Println 打印的输出结果是类似的。
	arr := [3]int{1, 2, 3}
	slice := [3]int{1, 2, 3}
	fmt.Println("arr", arr)     //arr [1 2 3]
	fmt.Println("slice", slice) //slice [1 2 3]

	//数据类型是否相同
	fmt.Println(reflect.TypeOf(arr) == reflect.TypeOf(slice)) // true
	//两个数组的元素是否相同
	fmt.Println(arr == slice) // false
}
