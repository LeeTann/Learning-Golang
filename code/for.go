package main

import "fmt"

func main() {
	// While loop
	// x := 2
	// for x <= 5 {
	// 	println(x)
	// 	x++
	// }

	// Traditional for loop
	// for i := 1; i <= 10; i += 2 {
	// 	fmt.Println(i)
	// }

	for y := -10; y <= 100; y++ {
		if y < 0 {
			println("Negative number")
			continue
		}

		if y%3 == 0 && y%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if y%3 == 0 {
			fmt.Println("Fizz")
		} else if y%5 == 0 {
			fmt.Println("Buzz")
		} else {
			println(y)
		}
	}

	//
	// var mySentence = "This is a sentence"

	// for index, letter := range mySentence {
	// 	fmt.Println("Index:", index, "Letter:", string(letter))
	// }
}
