package main

import "fmt"

type MyInt int

func (x MyInt) Double() MyInt {
	return x * 2
}

func main() {
	var x MyInt = 3
	fmt.Println(x.Double())
}
