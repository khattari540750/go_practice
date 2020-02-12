package main

import "fmt"


type Tag struct {
	Id int32
	X, Y, Z float32
}


func main() {

	mapT := map[int] Tag{}

	// var tag1  = Tag{1, 0.1, 0.2, 0.3}
	// var tag2  = Tag{1, 0.2, 0.3, 0.4}
	// var tag3  = Tag{1, 3.9, 1.3, 2.4}

	mapT[1] = Tag{1, 0.1, 0.2, 0.3}
	mapT[2] = Tag{2, 0.2, 0.3, 0.4}
	mapT[2] = Tag{2, 3.9, 1.3, 2.4}

	//fmt.Println(mapT)

	for key, _ := range mapT {
        fmt.Print(key)
        fmt.Print(" ")
    }
    fmt.Println()
}