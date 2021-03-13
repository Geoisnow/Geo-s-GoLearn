package lib2

import "fmt"

func Lib2Test()  {
	fmt.Println("this is lib2")
}
func init()  {
	fmt.Println("imported lib2 used init func..")
}