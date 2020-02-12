package main

import (
	"fmt"
	"./yaml"
)

func main() {
	k, e :=  yaml.ReadYamlFile("sample.yaml")
	if e != nil {
		return
	}
	fmt.Println(k.Users[0].Sex)
	fmt.Println(k.Users[1].Sex)
}