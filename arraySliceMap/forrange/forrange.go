package main

import "fmt"

func main() {
	//delcara o array, e o compilador define o tamanho de acordo com a quantidade de valores atribuidos a ele na inicializacao
	numeros := [...]int{1, 2, 3, 4, 5}

	//percorre o array e retorna indice e o valor em si
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i+1, numero)
	}

	//percorre o array ignorando o indice
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}

}
