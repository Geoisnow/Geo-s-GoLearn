package main

import (
	"fmt"
)

//延迟调用

func main() {

	defer func() {
		fmt.Println("first")
	}()

	defer func() {
		fmt.Println("second")
	}()
	fmt.Println("哈哈哈")
	//os.Exit(1) 使用os.Exit() 后所有defer都不执行
	a := 0
	defer func(i int) {
		fmt.Println("defer i=", i) //i=0 实参a的值 defer注册时通过值拷贝传递进去，a++不影响defer的输出结果
	}(a)
	a++
	return
}
//哈哈哈
//defer i= 0
//second
//first    从后往前执行，

/* 打开文件后要 defer fileName.Close()