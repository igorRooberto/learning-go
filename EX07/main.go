package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {

		switch {
		case i%2 == 0:
			fmt.Printf("%d é Par",i)
		default:
			fmt.Printf("%d é Impar",i)
		}
		fmt.Println()
	}

}