package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(Dividir(5,5))

}

func Dividir(a, b float64) (float64, error) {

	if b <= 0 {
		return 0, errors.New("Não é possível dividir por zero ou menor que zero;")
	}

	return a / b, nil
}