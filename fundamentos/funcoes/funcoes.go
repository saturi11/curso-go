package main

func main() {
	c := somar(5, 7)
	imprimir(c)
}

func imprimir(valor int) {
	println(valor)
}
func somar(a int, b int) int {
	return a + b
}
