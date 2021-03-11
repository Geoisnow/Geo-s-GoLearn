package main

import "fmt"

/*creat a new func
func Func_name(v_type1,v_type2)[r_v_type1,r_v_type2]{
	pass
}  */
func swap(a int,b int)(int,int){

return b,a
}

/*type 可以定义函数类型,其形参和返回值类型一样
 type FuncTypeName func()
1.var Func_name FuncTypeName
  Func_name = Exist_Func_name
  Func_name(para) 调用
2.Func_name := Exist_Func_name //自动推导函数类型
  Func_name(para) 调用
*/
//type 可以为已存在的类型起名

type Func_type1 func()  //没有参数也没有返回值的类型

func None_v_r()  {  // value return
	fmt.Println("im None_v_r")
}
func Have_v_r(a string) string{

	a = "the parameter is " +a
	return a
}

type Func_type2 func(string)string //有参数有返回值的类型

func main() {

	fmt.Println(swap(1,2))   // result  2 1
	var a Func_type1
	a = None_v_r
	a()			//	im None_v_r
	fmt.Printf("%T\n",a) //main.Func_type1

	var b Func_type2
	b = Have_v_r
	fmt.Println(b("haha")) //the parameter is haha
	fmt.Printf("%T\n",b) //main.Func_type2

	c := Have_v_r    		//自动推导函数类型
	fmt.Println(c("555")) //the parameter is 555
	fmt.Printf("%T\n",c) //func(string) string
}
