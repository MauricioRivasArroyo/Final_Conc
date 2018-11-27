package main

import (
	"fmt"
	"net"
	"bufio"
)

const local 	= "10.142.232.169:8001"
const next	= "10.142.232.185:8001"
const protocol	= "tcp"

func handle(conn net.Conn, mensajes chan string) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	mensajes<- msg[:len(msg)-1]
}

func receive(mensajes chan string) {
	ln, _ := net.Listen(protocol, local)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		handle(conn, mensajes)
	}
}

func send(remote, msg string) {
	conn, _ := net.Dial(protocol, remote)
	defer conn.Close()
	fmt.Fprintln(conn, msg)
}

func main() {
	fmt.Println("Esperando mensaje...")
	mensajes := make(chan string)
	go receive(mensajes)

	for msg := range mensajes {
		pos := 0
		min := rune(msg[0])
		for i, c := range msg {
			if c < min {
				min = c
				pos = i
			}
		}
		fmt.Printf("Número: %c\n", min)
		aux := []rune(msg)
		msg2 := string(append(aux[:pos], aux[pos+1:]...))

		fmt.Println("Reenviando mensaje...")
		send(next, msg2)
	}
}
