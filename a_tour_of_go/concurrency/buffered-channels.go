package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <-1
	ch <-2
	ch <-3

	cc := <-ch
	fmt.Println(cc)	

	cc = <-ch
	fmt.Println(cc)

	fmt.Println(<-ch)

	ch <- 4
	fmt.Println(<-ch)	
}