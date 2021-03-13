package lib1

import "fmt"

func Lib1Test()  {
	fmt.Println("this is lib1")
}
func init()  {
	fmt.Println("imported lib1 used init func..")
}