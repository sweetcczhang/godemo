package main

import "fmt"

var name int
func main()  {
	var a, b, c int
	d:=10
	a = 20
	b = 20
	c = a+b
	name = d
	fmt.Println(name)
	fmt.Print("Hello world;")
	sum(a, c)
}
func sum(a int,b int)  {
	var c int
	c = a + b
	fmt.Printf("sumc()函数中 a = %d\n", c)
}
