package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sequential Loop:")
	fmt.Println("----------------")
	sequential()
	fmt.Println()
	fmt.Println("Concurrent Loop:")
	fmt.Println("----------------")
	concurrent()
}