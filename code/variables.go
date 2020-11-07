package main

import "fmt"

func main() {
	var name2 string = "Rebs"
	var number = 260
	number = number + 5
	name := "Lee" // shorthand way is called typed inference
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(number)
	fmt.Printf("%T", number)
}
