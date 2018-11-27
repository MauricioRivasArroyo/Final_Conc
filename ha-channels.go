package main

import (
	"fmt"
)

var ch chan string

func p() {
	ch<- "Hola esto es un ejemplo"
}

func main() {
	ch = make(chan string)
	go p()
	msg := <-ch
	fmt.Println(msg)
}