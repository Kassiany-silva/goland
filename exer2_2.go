/*Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
Demonstre estes valores.*/
package main

import "fmt"

var a = 10

func main() {
	b := (50 == a)
	c := (10 == 10)
	d := (10 <= 30)
	e := (10 >= 20)

	fmt.Printf("%v\n%v\n%v\n%v\n", b, c, d, e)
}
