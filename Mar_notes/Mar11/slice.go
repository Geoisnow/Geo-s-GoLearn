package main

import "fmt"

func main()  {
	//数组定义 var arrName [len]arr_type
	//切片定义 var sliceName []type     不写元素个数的叫切片，写的叫数组
	a := [4]int{1,2,3,4} //确定大小后不能添加
	//a = append(a,5) err
	var b []int
	b = append(b,1,2,3,4)

	c := make([]int,2)
	c[0] = 1
	c[1] = 2
	c = append(c,3)		//while cap<1024 会扩大上次的2倍,cap>1024 扩大1/4
	fmt.Println(a,b,c)  		//[1 2 3 4] [1 2 3 4] [1 2 3]
	fmt.Println(a[2:],b[:2],c) //切片 [3 4] [1 2] [1 2 3]
	/*
	for i := range a{		//遍历输出
		fmt.Println(a[i])
	}	*/
	d := make([]int,5)
	d[0] = 1
	d[1] = 2
	fmt.Println(d)  // [1 0 0 0 0] 默认为0
	d = append(d,2,3,4)		//append 的长度后面添加
	fmt.Println(d)  //[1 0 0 0 0 2 3 4]
	/* 拷贝 numbers 的内容到 numbers1 */
	//copy(numbers1,numbers)
	e := make([]int,2)
	e[0] = 6
	e[1] = 7
	copy(e,d)  	//根据长度复制
	fmt.Println(e,d)
}
