/*
	Usando uma literal composta:

Crie uma slice de tipo int
Atribua 10 valores a ela
Utilize range e demonstre os valores do
Utilizando format printing, demonstre o tipo do slice.!
*/
package main

import "fmt"

func main() {

	slice := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
}
