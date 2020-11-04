package main

import "fmt"

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i)
	// }

	var mySentence = "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", string(letter))
	}
}
