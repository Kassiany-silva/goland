/*
Crie um map com key tipo string e value tipo []string.
Key deve conter nomes no formato sobrenome_nome
Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.
*/
package main

import "fmt"

func main() {

	mapa := map[string][]string{
		"teste": []string{
			"oi", "olá", "hi"},

		"bom dia ": []string{
			"good morning", "afternoon", "evening"},

		"lindo": []string{
			"dia ", "sorriso", "bolo"},
	}
	for k, v := range mapa {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}

	}

}
