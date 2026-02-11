package main

import "fmt"

func main() {
	dividir(10, 5)
	dividir(10, 0)
}

func dividir(a, b int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado:", r)
		}
	}()

	if b == 0 {
		panic("não é possível dividir por zero")
	}

	fmt.Println(a / b)
}