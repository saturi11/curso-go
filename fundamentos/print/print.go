package main

import "fmt"

func main() {
	//nao cira separacao de linha, escreve tudo na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//cria uma nova quebra de linha a cada comando
	fmt.Println(" Nova")
	fmt.Println("linha.")
	x := 3.141516

	// fmt.Println("O valor de x é " + x), nao consegue concatenar float com string no mesmo comando

	//retorna uma string atravez de um valor passado
	xs := fmt.Sprint(x)

	//transfoma em string e concatena no mesmo comando
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	//formata a string, %f significa que a variavel passada e um float, podeia ser %s indicando que a variavel e uma string ou %v que imprime todos os tipos
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
