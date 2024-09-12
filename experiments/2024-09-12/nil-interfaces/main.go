package main

import "fmt"

type MyInt int

type Foo interface {
	CheckNil() bool
}

func (v *MyInt) CheckNil() bool {
	return v == nil
}

func evaluate(v Foo) {
	fmt.Printf("Is the interface nil? %t\n", v == nil)
	fmt.Printf("CheckNil says:        %t\v", v.CheckNil())
	fmt.Println()
}

func main() {
	var x MyInt = 3
	var y *MyInt

	fmt.Println("Assigning non-nil pointer to z")
	var z Foo = &x
	evaluate(z)

	fmt.Println("Assigning nil pointer to z")
	z = y
	evaluate(z)
}
