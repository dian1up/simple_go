package main

import (
	"fmt"
)

func main() {
	fmt.Println(count("Hai goolang"))
}

func count(x string) string {
	for index := 0; index < len(x); index++ {
		return x[index]
	}
}
