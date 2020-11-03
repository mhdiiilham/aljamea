package main

import (
	"fmt"
)

func main() {
	jamaah := newJamaah()
	randomIndex := randomIndex(len(jamaah))
	fmt.Println(randomIndex)
}
