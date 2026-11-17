package main

import "fmt"

// For循环
func main() {

	//for 是 Go 中唯一的循环结构。这里会展示 for 循环的一些基本使用方式。

	//1 最基础的方式，单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//2 经典的初始/条件/后续 for 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//3 不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。
	for {
		fmt.Println("loop")
		break
	}

	//4 使用 continue 直接进入下一次循环
	for n := 0; n <= 5; n++ {
		// 如果n是偶数，就进入下一轮循环
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
