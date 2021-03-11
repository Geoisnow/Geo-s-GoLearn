package main

import "fmt"

func main() {
	/*creat a struct
	type struct_name struct {
		feildName1 feildType
		feildName2 feildType
		feildName3 feildType
	} 	*/
	type struct1 struct {
		num int
		name string
		exam bool  //默认false
	}
	//几种构造方法
	a := struct1{1,"张三",true} //must be 3 vaules,if not err : too few values in struct1 literal
	b := struct1{num: 2,name: "李四",exam: false}  //可以低于3个value

	var c struct1
	c.num = 3
	c.name = "王五"
	c.exam = true

	d := &struct1{
		num: 4,
		name: "马六",
	}
	d.exam = true

	fmt.Println(a,b,c,d) //{1 张三 true} {2 李四 false} {3 王五 true} &{4 马六 true}
	fmt.Println(a.name)  //张三

	type strcut2 struct {
		struct1
		comment string
	}
	e := strcut2{a,"very good"}
	fmt.Println(e)  //{{1 张三 true} very good}


}