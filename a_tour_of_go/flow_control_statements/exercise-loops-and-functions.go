package main

import "fmt"

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		zold := z
		z = z - (z*z-x)/(2*z)
		if zold-z == 0 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(sqrt(3))
}
