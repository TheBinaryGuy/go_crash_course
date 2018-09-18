package main

import (
	"fmt"
)

func main() {
	// x := 15
	// y := 10

	// if x < y {
	// 	fmt.Printf("%d is less than %d\n", x, y)
	// } else {
	// 	fmt.Printf("%d is less than %d\n\n", y, x)
	// }

	color := "red"
	// if color == "red" {
	// 	fmt.Println("color is red")
	// } else if color == "green" {
	// 	fmt.Println("color is green")
	// } else {
	// 	fmt.Println("color is neither green nor red")
	// }

	switch color {
	case "red":
		fmt.Println("color is red")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("color is neither green nor red")
	}
}
