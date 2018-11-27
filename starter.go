package main

import (
	"fmt"
	"net"
)

const next	= "localhost:8060"
const protocol	= "tcp"

func send(remote, msg string) {
	conn, _ := net.Dial(protocol, remote)
	defer conn.Close()
	fmt.Fprintln(conn, msg)
}

func main() {
	send(next, "4918375602")
}
