package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 exibe o texto de toda linha que aparece mais de
// uma vez na entrada-padrão, precedida por sua contagem.
func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Println(counts)

	for input.Scan() {
		counts[input.Text()]++
	}

	//NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}

}
