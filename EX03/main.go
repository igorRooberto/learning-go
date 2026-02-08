package main

import "fmt"

func main() {

	fmt.Println("Arara é palíndromo?", EhPalidromo("arara"))

	fmt.Println("Bob é palíndromo?", EhPalidromo("bob"))

}

func EhPalidromo(Word string) bool {

	tamanho := len(Word)

	for i := 0; i < tamanho/2; i++ {

		if Word[i] != Word[tamanho-1-i] {
			return false
		}
	}
	return true
}