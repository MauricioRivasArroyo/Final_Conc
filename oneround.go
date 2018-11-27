package main

import (
	"bufio"
	"fmt"
	"net"
)

type Plan struct {
	general string
	plan    rune
}

const local = "10.142.232.169:8001"
const protocol = "tcp"
const nombre = "martin"
var addresses []string = []string{"10.142.232.185:8001",
                                  "10.142.232.186:8001",
                                  "10.142.232.189:8001",
                                  "10.142.232.183:8001"}

func receiveAll(numGenerals int, plans chan Plan) {
	ln, _ := net.Listen(protocol, local)
	defer ln.Close()
	for i := 0; i < numGenerals; i++ {
		conn, _ := ln.Accept()
		go receive(conn, plans)
	}
}

func receive(conn net.Conn, plans chan Plan) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	size := len(msg)
	plans<- Plan{ msg[:size-2], rune(msg[size-2]) }
}

func sendAll(plan rune) {
	for _, remote := range addresses {
		send(remote, plan)
	}
}

func send(remote string, plan rune) {
	conn, _ := net.Dial(protocol, remote)
	defer conn.Close()
	fmt.Fprintf(conn, "%s%c\n", nombre, plan)
}

func main() {
	plans := make(chan Plan)
	numGenerals := len(addresses)
	go receiveAll(numGenerals, plans)
	var x string
	fmt.Println("Presione enter para enviar plan a todos...")
	fmt.Scanf("%s\n", &x)
	fmt.Println(x)
	plan := 'A'
	sendAll(plan)
	cont := 0
	if plan == 'A' {
		cont++
	}
	for i := 0; i < numGenerals; i++ {
		plan := <-plans
		fmt.Printf("%s: %c\n", plan.general, plan.plan)
		if plan.plan == 'A' {
			cont++
		}
	}
	if cont > (numGenerals+1) / 2 {
		fmt.Println("ATACAMOS!")
	} else {
		fmt.Println("NOS RETIRAMOS!")
	}
}
