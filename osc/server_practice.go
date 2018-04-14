package main

import (
	"github.com/hypebeast/go-osc/osc"
)

func main() {
	addr := "127.0.0.1:10002"
	server := &osc.Server{Addr: addr}

	server.Handle("/comm/ball/color", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server.ListenAndServe()


}