package main

import "fmt"

func main() {
	var point = 2

	if point == 10 {
		fmt.Println(point)
	} else if point < 10 {
		fmt.Println("kurang dari 10")
	} else if point > 10 {
		fmt.Println("lebih dari 10")
	} else {
		fmt.Println("print ini tidak mungkin muncul")
	}
}
