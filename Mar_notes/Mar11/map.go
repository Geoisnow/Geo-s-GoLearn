package main

import "fmt"

func main()  {
	//creat map
	//var:= map[keyType]valueType{"k":v,}
	a := map[string]int{"a":1,"b":2,"c":3}
	a["d"]=4
	fmt.Println(a)

	//make函数创建
	//make(map[k]t)
	//make(map[k]t,len)
	//a = {"a":1,"b":2,"c":3}
	b :=make(map[string]int)
	b["a"]=1
	b["b"]=2
	b["c"]=3
	delete(b,"c") //map删除
	b["a"]=6//修改
	fmt.Println(b)


	for k,v := range a{
		fmt.Println(k,v)
	}

}
/*运行结果
map[a:1 b:2 c:3 d:4]
map[a:6 b:2]
a 1
b 2
c 3
d 4

 */