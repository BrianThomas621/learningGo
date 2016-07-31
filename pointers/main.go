package main

import "fmt"

func pass_intvalue(z int) {
	z = 0
}

func pass_pointervalue(z *int) {
	*z = 0
}

func main() {
	x := 5
	fmt.Println("Value of x: ", x)
	fmt.Println("Address of x: ", &x)
	pass_intvalue(x)
	fmt.Println("Not passing a memory address > Value of x:", x)
	pass_pointervalue(&x)
	fmt.Println("After passing a memory address > Value of x:", x)
}
