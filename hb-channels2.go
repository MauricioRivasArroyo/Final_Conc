package main

import (
	"fmt"
)

var ch chan rune

func p() {
	msg := "Hola, mundo"
	for _, c := range msg {
		ch<- c
	}
	close(ch)
}

func main() {
	ch = make(chan rune)
	go p()
	for v := range ch {
		fmt.Printf("'%c' ", v)
	}
}
