package main

import (
	"fmt"
	"math"
)

func main() {
	//cria uma constante declarando nome e tipo
	const PI float64 = 3.1415

	//cria uma var, nao precisa declarar tipo se for atribuido valor
	var raio = 3.2

	//declara e inicializa uma variavel de forma reduzida
	area := PI * math.Pow(raio, 2)

	fmt.Println("area do circulo:%s", area)
	//declara varias constantes de uma vez
	const (
		a = 1
		b = 2
	)
	//declara varias variaveis de uma vez
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
	var e, f bool = true, false
	fmt.Println(e, f)
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
