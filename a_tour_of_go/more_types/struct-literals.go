package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 3, X: 4}
	v3 = Vertex{}
	p  = &Vertex{5, 6}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
