package main

import (
	"fmt"
)

func concurrent() {
	go callLoop("Hello, Round Earth!", 10)
	go callLoop("Hello, Flat Earth!", 10)
	fmt.Scanln()
}