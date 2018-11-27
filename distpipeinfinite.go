package main

import (
	"fmt"
	"net"
	"bufio"
)

const local = "localhost:8060"
const next	= "localhost:8070"
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
	fmt.Fprintf(conn, "%s,luis\n", msg)
}

func main() {
	fmt.Println("Esperando mensaje...")
	mensajes := make(chan string)
	go receive(mensajes)

	for msg := range mensajes {
		fmt.Println("Mensaje recibido: ", msg)

		fmt.Println("Reenviando mensaje...")
		send(next, msg)
	}
}
