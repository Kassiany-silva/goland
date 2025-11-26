package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}
	uorigem := os.Args[len(os.Args)-1]
	vorigem := os.Args[1 : len(os.Args)-1]
	var udestino string
	if uorigem == "c" {
		udestino = "f"

	} else if uorigem == "km" {
		udestino = "mi"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", udestino)
		os.Exit(1)

	}
	for i, v := range vorigem {
		vorigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)

		}
		var vdestino float64
		if uorigem == "c" {
			vdestino = vorigem*1.8 + 32

		} else {
			vdestino = vorigem / 160934
		}
		fmt.Printf("%.2f %s = %.2f %s \n", vorigem, uorigem, vdestino, udestino)
	}

}
/*exemplo de teste 32 27.4 -3 0 c*/