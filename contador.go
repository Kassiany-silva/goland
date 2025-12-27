package main

import (
	"fmt"
	"os"
)

func main() {
	palavras := os.Args[1:]
	estatisticas := colherEstatisticas(palavras)
	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estatisticas := make(map[string]int)
	for _, palavra := range palavras {
		if len(palavra) > 0 {
			inicial := string(palavra[0])
			estatisticas[inicial]++
		}
	}
	return estatisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")
	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}
