package main

import (
	"fmt"
	"math/rand"
	"time"
)

// cria uma variavel que vai ser inicializada junto com o programa
func numeroAleatorio() int {
	//gera um numero aleatorio pegando a hora atual e pega o nanosegundo atual da execucao para inicializar a variavel
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
