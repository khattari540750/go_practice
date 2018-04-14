package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["test"] = Vertex{123.5, -8811.2}
	m["test2"] = Vertex{221.1, -2321.4}
	fmt.Println(m["test2"])
}
