package main

import "fmt"

type Hero struct {
	Name string
	Hp int
	Level int
}
/*
func (this Hero) Show()  {  //this 类似python类中的self
	//this是调用该方法的对象的一个副本，修改不了该对象
	fmt.Println("Name = ",this.Name)
	fmt.Println("Hp = ",this.Hp)
	fmt.Println("Level = ",this.Level)
}

func (self Hero)GetName()string{
	return self.Name
}

func (this Hero)SetName(NewName string)  {
	this.Name = NewName
}

func main() {
	hero := Hero{"托儿索",250,1}
	hero.Show()
	hero.SetName("儿童劫")
	fmt.Println(hero.GetName())
}

Name =  托儿索
Hp =  250
Level =  1
托儿索
 */
func (this *Hero) Show()  {
	fmt.Println("Name = ",this.Name)
	fmt.Println("Hp = ",this.Hp)
	fmt.Println("Level = ",this.Level)
}

func (self *Hero)GetName()string{
	return self.Name
}

func (this *Hero)SetName(NewName string)  {
	this.Name = NewName
}

func main() {
	hero := Hero{"托儿索",250,1}
	hero.Show()
	hero.SetName("儿童劫")
	fmt.Println(hero.GetName())
}
/*
Name =  托儿索
Hp =  250
Level =  1
儿童劫
*/