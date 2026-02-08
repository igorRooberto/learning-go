package main

import "fmt"

func main() {

	var numeroTabuada int

	fmt.Println("Digite a Tabuada que você deseja: ")
	fmt.Scan(&numeroTabuada)

	for i := 1; i <= 10; i++{
		fmt.Printf("%d X %d = %d\n", numeroTabuada, i , numeroTabuada*i)
	}

}