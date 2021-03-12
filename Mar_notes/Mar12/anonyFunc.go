package main

import "fmt"


func add(a,b int) int{
	return a+b
}

func sub(a,b int)int {
	return a-b
}

type Operate func(int,int)int  //定义一个函数类型，输入的是两个 int 类型，返回位是一个int类型

func do(function Operate,a,b int) int{ //主义一个函数，第1一个参数是函数类型 Operate
	return function(a,b)
}

func wrap(op string) func(int,int)int{
	switch op {
	case "add":
		return func(a,b int)int {
			return a+b
		}
	case "sub":
		return func(a,b int)int {
			return a-b
		}
	default:
		return nil
	}
}

func main() {
	a := do(add, 1, 2)
	b := do(sub, 4, 2)

	cadd := add //函数可以复制给变量

	c := cadd(2, 3)
	fmt.Println(a, b, c) //3 2 5

	//匿名函数
	var sum = func(a, b int) int { return a + b }
	d := sum(4, 5)
	fmt.Println(d)



	wrap_add := wrap("add")
	e:=wrap_add(1,2)
	fmt.Println(e)  //3

}