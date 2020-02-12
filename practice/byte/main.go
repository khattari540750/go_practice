package main

import(
	"strings"
	"fmt"
	//"bytes"
)

func main() {
	vec := "POS,0x3304,x,y,z\n"
	sList := strings.Split(vec, ",")
	for _, s := range sList {
		fmt.Println(s)
	}
	println(len(sList))
}