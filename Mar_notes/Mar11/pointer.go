package main

import "fmt"
func main() {
	//指针类型用于传递地址, 而不是传递值,值大时效率低，指针用来提高效率
	//var ptr_name *ptrType 创建指针
	var a int= 20
	var ptr1 *int
	ptr1 = &a //整数地址赋值给指针

	fmt.Printf("a变量的地址是: %x\nptr1的地址是: %x\n",&a,ptr1)
	//a变量的地址是: c00000a090
	//ptr1的地址是: c00000a090

	fmt.Printf("a的值为:%d\nptr1指向的值为:%d\n",a,*ptr1)
	//a的值为:20
	//ptr1指向的值为:20

	//指针数组
	b:= []int{10,100,200}
	const n = 3
	//fmt.Println(len(b))
	var ptr2 [n]*int
	for i := 0; i < n; i++ {
		ptr2[i] = &b[i] //整数地址赋值给指针数组

	}
	for i := 0; i < n; i++ {
		fmt.Printf("b[%d]的地址为：%x,值为:%d\n",i,&b[i],b[i])
		fmt.Printf("ptr2[%d]的地址为：%x,指向的值为:%d\n",i,ptr2[i],*ptr2[i])
	}
	/*
	b[0]的地址为：c000012360,值为:10
	ptr2[0]的地址为：c000012360,指向的值值为:10
	b[1]的地址为：c000012368,值为:100
	ptr2[1]的地址为：c000012368,指向的值值为:100
	b[2]的地址为：c000012370,值为:200
	ptr2[2]的地址为：c000012370,指向的值值为:200
	*/

	// Pointer's pointer
	c := 66
	var ptr3 *int
	var ptr4 **int
	ptr3 = &c
	ptr4 = &ptr3
	fmt.Printf("c的值为:%d，地址为:%x\nptr3指向的值为%d:,地址为%x\n",c,&c,*ptr3,ptr3)
	fmt.Printf("ptr3指向的的值为:%d，地址为:%x\nptr4指向的值为:%x,地址为%x\n",*ptr3,ptr3,*ptr4,ptr4)
	/*
	c的值为:66，地址为:c00000a098
	ptr3指向的值为66:,地址为c00000a098
	ptr3指向的的值为:66，地址为:c00000a098
	ptr4指向的值为:c00000a098,地址为c000006030
	 */

	type struct1 struct {
		name string
		age int
	}
	r := struct1{"张三",10}
	p := &r
	fmt.Printf("r的地址为%x：,p的地址为%x",&r,p)
	fmt.Println(*p)
	var aa = [3]int{1,2,3}
	pp := &aa
	fmt.Println(pp)
	//r的地址为&{e5bca0e4b889 a}：,p的地址为&{e5bca0e4b889 a}{张三 10}
	//&[1 2 3]
	}

