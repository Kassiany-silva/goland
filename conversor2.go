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
	} else if uorigem == "f" {
		udestino = "c"
	} else if uorigem == "km" {
		udestino = "mi"
	} else if uorigem == "mi" {
		udestino = "km"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!\n", uorigem) // Corrigido: uorigem em vez de udestino
		os.Exit(1)
	}

	for i, v := range vorigem {
		valor, err := strconv.ParseFloat(v, 64) // Corrigido: usar nome diferente
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var vdestino float64
		if uorigem == "c" {
			vdestino = valor*1.8 + 32
		} else if uorigem == "f" {
			vdestino = (valor - 32) / 1.8
		} else if uorigem == "km" {
			vdestino = valor * 0.62137
		} else if uorigem == "mi" {
			vdestino = valor / 0.62137 // Corrigido: divisão pelo fator correto
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valor, uorigem, vdestino, udestino)
	}
}
