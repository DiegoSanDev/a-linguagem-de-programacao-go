// Echo1 exibe seus argumentos de linha de comando.clear
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Index:", i, "Programa: "+s)
	}
	fmt.Println("Time difference:", time.Now().Nanosecond()-start.Nanosecond())
}
