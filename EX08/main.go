package main

import "fmt"

func main() {

	var celsius float64

	fmt.Println("Insira uma temperatura em C° : ")
	fmt.Scan(&celsius)

	fahrenheit := celsius * 1.8 + 32

	fmt.Println("A temperatura em fahrenheit = ",fahrenheit)

}