// Echo1 exibe seus argumentos de linha de comando.clear
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Time difference:", time.Now().Nanosecond()-start.Nanosecond())
}
