package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Printf("%v ", i)
		i = i + 1
	}

	fmt.Println()

	for j := 7; j <= 9; j++ {
		fmt.Printf("%v ", j)
	}

	fmt.Println()
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Printf("%v ", n)
	}
}
