package main

import (
	"curso-go/formas"
	"fmt"
)

func main() {

	meuQuadrado := formas.Quadrado{Lado : 10}
	
	fmt.Println("A área do Quadrado é de: ")
	Medir(meuQuadrado)
}

func Medir(g formas.Geometria){
	fmt.Println("Calculando a área...")
    fmt.Println("Resultado:", g.Area())
}