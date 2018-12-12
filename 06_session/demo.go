package main

import "fmt"

func main() {
	const c1 string = "hello"
	const c2  = "world"

	const n1 int8 =100
	fmt.Printf("%T\n",n1)
	const n2 = 100
	fmt.Printf("%T\n",n2)
}
