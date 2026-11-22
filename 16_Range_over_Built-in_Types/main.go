package main

import (
	"fmt"
)

// 遍历内置类型
// range 用于遍历各种内置数据结构中的元素。下面来看看如何使用 range 遍历我们已经学过的某些数据结构。
func main() {

	//1 这里我们使用 range 来对切片中的数字求和。数组也可以这样使用。
	var nums = []int{2, 3, 4}
	var sum = 0
	for _, num := range nums {
		sum = sum + num
	}

	fmt.Println(sum) //9

	//2 数组和切片上的范围为每个条目同时提供了索引和值。
	//  在上文中，我们不需要索引，因此我们用空白标识符 _ 将其忽略。但有时我们确实需要索引。

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) //index: 1
		}
	}

	//3 映射中的范围遍历键/值 对。
	//创建并初始化map,map[键的类型]值的类型
	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		//打印每个键值对
		//Printf 按格式输出：%s 是字符串占位符，\n 是换行
		fmt.Printf("%s -> %s\n", k, v)
	}

	//4 range 也可以仅遍历映射中的键。

	for k := range kvs {
		fmt.Println("key:", k)
	}

	//5 字符串上的range迭代遍历Unicode码位。
	//  第一个值是字符的起始字节索引，第二个值是字符本身。更多详情请参见“字符串与字符”。

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
*/
