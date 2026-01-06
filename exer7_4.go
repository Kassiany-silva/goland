/*
Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import "fmt"

func main() {

	x := [][]string{
		[]string{
			"ana",
			"leite",
			"ler"},
		[]string{
			"pedro",
			"henrique",
			"comer"},
		[]string{
			"livia",
			"dantas",
			"correr"},
	}
	for _, v := range x {
		fmt.Println(v)
	}
	for _, v := range x {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}
