package main

import (
	"fmt"
)

func callLoop(message string, rounds int) {
	for i := 0; i < rounds; i++ {
		fmt.Println(message)
	}
}