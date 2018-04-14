package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Test":   {123.1, -123.1},
	"Google": {33.54, -222.44},
}

func main() {
	fmt.Println(m)
}
