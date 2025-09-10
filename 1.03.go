package main

import "fmt"



func method1(){
	a:=1
	var b int = 2
	var c int 
	fmt.Println(a,b,c)
}
func method2()(a int,b string){
	a=10
	b="hello"
	return a,b
}

func main() {
	fmt.Println("Hello, world!")
	method1()
	var a,b = method2();
	fmt.Println("method2()",b)
	fmt.Println("a=",a)
}