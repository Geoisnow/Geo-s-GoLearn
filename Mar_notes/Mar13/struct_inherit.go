package main

import "fmt"

type Man struct {
	name string
	sex string
}

func (this *Man)walk()  {
	fmt.Println("Man.walk()...")
}
func (this *Man)sing()  {
	fmt.Println("Man.sing()...")
}
//=====
type SuperMan struct {
	Man		//SM类继承了Man类的方法
	body string
}

func (this *SuperMan)walk()  {
	fmt.Println("SuperMan.walk()...")
}
//定义子类新方法
func (this *SuperMan)fly(){
	fmt.Println("SuperMan.flying()...")
}
func main() {

	man := Man{"张三","男"}
	man.walk()
	man.sing()
	//superm := SuperMan{Man{"超人","男"},"强壮"}
	var superm SuperMan
	superm.name = "超人"
	superm.sex = "男"
	superm.body = "强壮"
	superm.walk() 	//子类定义的新方法
	superm.sing()	//父类的方法
	superm.fly()
/*
   Man.walk()...
   Man.sing()...
   SuperMan.walk()...
   Man.sing()...
   SuperMan.flying()...
 */



}
