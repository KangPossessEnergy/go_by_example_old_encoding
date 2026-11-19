package main

import (
	"fmt"
	"time"
)

// switch 是多分支情况时快捷的条件语句。
func main() {

	//1 一个基本的 switch。
	i := 2
	fmt.Println("write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//2 在同一个 case 语句中，你可以使用逗号来分隔多个表达式。 在这个例子中，我们还使用了可选的 default 分支
	//time.Now().Weekday() 获取当前日期对应的星期几
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //如果今天说周六或周日
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//3 不带表达式的 switch 是实现 if/else 逻辑的另一种方式。 这里还展示了 case 表达式也可以不使用常量。

	t := time.Now() //Go 标准库 time 包里的一个函数，作用是返回当前时间。
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//4 类型开关 (type switch) 比较类型而非值。可以用来发现一个接口值的类型。
	//  在这个例子中，变量 t 在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) { //定义匿名函数
		switch t := i.(type) { //i.(type)：特殊语法，只能用在switch里,表示"判断i的动态类型""
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t) //Printf：第一个参数是格式串，里面的 %T（类型）、%d（整数）、%s（字符串）、\n（换行）都会被解析，后面的参数按顺序填入。
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
