package main

import "fmt"

const myConstant string = "constant"

func main() {
	var myInitialVar = "initial"
	fmt.Println(myInitialVar)

	myInitialVar = "overwrite initial"
	fmt.Println(myInitialVar)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e string
	fmt.Println(e == "")


	var z float32
	fmt.Println(z)

	fmt.Println(myConstant)

}
