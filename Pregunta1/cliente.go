package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var mat string

func dibujarMatriz(mat string) {
	fmt.Println()
	for i := 0; i < len(mat)-1; i += 3 {
		fmt.Printf(" %s | %s | %s\n", string(mat[i]), string(mat[i+1]), string(mat[i+2]))
		if i < 6 {
			fmt.Println("-----------")
		}
	}
	fmt.Println()
}

func enviar() {
	conn, _ := net.Dial("tcp", "localhost:8001")
	fmt.Print("Posicion (0-8): ")
	gi, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	defer conn.Close()
	fmt.Fprintln(conn, gi)
}

func main() {
	fmt.Println("Jugador: Elegir entre las posiciones desde 0 hasta 8")
	ln, _ := net.Listen("tcp", "localhost:8002")
	for {
		enviar()
		conn, _ := ln.Accept()
		r, _ := bufio.NewReader(conn).ReadString('\n')
		mat = r
		if r[0] == '!' {
			fmt.Println(r)
			break
		}
		dibujarMatriz(mat)
	}

	ln.Close()
}
