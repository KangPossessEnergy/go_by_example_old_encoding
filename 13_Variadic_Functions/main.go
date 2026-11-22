package main

import "fmt"

// 1 这个函数接受任意数量的int作为参数
func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	//用for range遍历这个切片,把所有数累加到total
	for _, num := range nums {
		total += num
		// total= total+num
	}
	fmt.Println(total)
}

// 变参函数
// 在调用时可以传递任意数量的参数。 例如，fmt.Println 就是一个常见的变参函数。
func main() {

	//变参函数使用常规的调用方式，传入独立的参数
	sum(1, 2)
	sum(1, 2, 3)

	//如果你有一个含有多个值的slice，想把它们作为参数使用，你需要这样调用 func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
