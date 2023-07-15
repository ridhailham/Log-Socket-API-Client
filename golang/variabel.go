package main

import "fmt"

func main() {
	var firstName string = "Ridha Ilham"
	const lastName = "Adi Setyawan"

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println("Halooo", firstName)

	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 100
	fmt.Println(*numberB)

	r := "Ridha Ilham Adi Setyawan"
	fmt.Println(r)

	i := 1
	for {
		fmt.Println(i)

		if i == 100 {
			fmt.Println("selesai")
			break
		}

		i++
	}

}
