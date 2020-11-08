package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Println("you can ride")
	} else if age >= 14 {
		fmt.Println("you can ride with parent")
	} else {
		fmt.Println("you cannot ride")
		fmt.Printf("Wait %d years", 18-age)
	}

	// Not DRY code
	// if age >= 18 {
	// 	fmt.Println("You can ride alone!")
	// }
	// if age >= 14 && age < 18 {
	// 	fmt.Println("You can ride with a parent")
	// }
	// if age < 14 {
	// 	fmt.Println("You cannot ride")
	// }

}
