package main

import (
	"fmt"
	"sort"
)

/*
Pec¸a ao usuario para digitar dez valores num ´ ericos e ordene por ordem crescente esses ´
valores, guardando-os num vetor. Ordene o valor assim que ele for digitado. Mostre ao
final na tela os valores em ordem.
*/

func main() {

	var vetor [10]int

	slice := []int{}

	var entrada int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um valor: ")
		fmt.Scan(&entrada)

		vetor[i] = entrada
		slice = append(slice, vetor[i])

	}

	sort.Ints(slice)

	fmt.Println("Valores ordenados: ", slice)


}
