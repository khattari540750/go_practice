package main

import (
	"fmt"
)


func main() {
	fmt.Println("process start.")
	defer printEnd()
	for {

	}
}

func printEnd() {
	fmt.Println("end")
}