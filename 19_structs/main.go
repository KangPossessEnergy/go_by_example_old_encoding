package main

import "fmt"

// 1 这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

// 2 newPerson 使用给定的名字构造一个新的 person 结构体
func newPerson(name string) *person {

	//3 您可以安全地返回指向局部变量的指针， 因为局部变量将在函数的作用域中继续存在。
	p := person{name: name}
	p.age = 42
	return &p
}

// Go 的结构体(struct) 是带类型的字段(fields)集合。 这在组织数据时非常有用。
func main() {
	//4 使用这个语法创建新的结构体元素
	fmt.Println(person{"kkdw", 18}) //{kkdw 18}

	//5 你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "kkdw", age: 18}) //{kkdw 18}

	//6 省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"}) //{Fred 0}

	//7 & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40}) //&{Ann 40}

	//8 在构造函数中封装创建新的结构实例是一种习惯用法
	fmt.Println(newPerson("Jon")) //&{Jon 42}

	//9 使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) //Sean

	//10 也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age) //50

	//11 结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age) //51

}
