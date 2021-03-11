package main

import "fmt"

func main() {
	//数组的创建
	var a = [3]int{1,2,3}
	b := [3]int{2,3,4}
	//fmt.Println(a)

	for i := range a{
		fmt.Println(a[i])
	}

	for _,v := range b{
		fmt.Println(v)
	}
	//[...]int{} 为不指定长度
	c := [...]int{1,2,3,4,5}

	fmt.Println(c)
	//切片运用类似python
	//通过make创建
	d := make([]int,2)
	fmt.Println("d为：",d,"长度为",len(d)," 容量为",cap(d))
	d = append(d, 1) //c = append(c, 6) 会报错，初始为几个就为几个
	fmt.Println("d为：",d,"长度为",len(d)," 容量为",cap(d))
	d = append(d, 1)
	fmt.Println("d为：",d,"长度为",len(d)," 容量为",cap(d))
	d = append(d, 1)
	fmt.Println("d为：",d,"长度为",len(d)," 容量为",cap(d))

	//	1
	//	2
	//	3
	//	2
	//	3
	//	4
	//  [1 2 3 4 5]
	//	d为： [0 0] 长度为 2  容量为 2
	//	d为： [0 0 1] 长度为 3  容量为 4
	//	d为： [0 0 1 1] 长度为 4  容量为 4
	//	d为： [0 0 1 1 1] 长度为 5  容量为 8  make创建的数组更灵活

	}
