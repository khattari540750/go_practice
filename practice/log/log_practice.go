package main

import (
	"log"
	"os"
	"io"
)

func main() {
	logfile, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()

	// log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetOutput(io.MultiWriter(logfile))

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Log!!")
}
