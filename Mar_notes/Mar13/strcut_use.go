package main

import "fmt"

type Man struct {
	name string
	age int
}

func (this Man)walk()  {
	fmt.Printf("%s is walking\n",this.name)
}
func (this Man)sing() {
	fmt.Printf("%s is singing\n",this.name)
}
func (this Man)info()  {
	fmt.Printf("%s is %d years old",this.name,this.age)
}
func main() {
	man := Man{"张三",16}
	man.walk()
	man.sing()
	man.info()
}
