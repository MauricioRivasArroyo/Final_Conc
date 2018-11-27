package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var ln net.Listener
var conn net.Conn
var err error
var mat []rune

func dibujarMatriz(mat []rune) {
	fmt.Println()
	for i := 0; i < len(mat); i += 3 {
		fmt.Printf(" %s | %s | %s\n", string(mat[i]), string(mat[i+1]), string(mat[i+2]))
		if i < 6 {
			fmt.Println("-----------")
		}
	}
	fmt.Println()
}

func enviar(reply string) {
	conn, _ := net.Dial("tcp", "localhost:8002")
	defer conn.Close()
	fmt.Fprintln(conn, reply)
}

func handlError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		ln.Close()
	}
}

func cantEspacios(mat []rune) int {
	cont := 0
	for i := 0; i < len(mat); i++ {
		if mat[i] == ' ' {
			cont++
		}
	}
	return cont
}

func validarSiHayGanador(mat []rune) string {
	//horizontales
	if mat[0] == mat[1] && mat[1] == mat[2] && (mat[0] == 'x' || mat[0] == 'o') {
		if mat[0] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	if mat[3] == mat[4] && mat[4] == mat[5] && (mat[3] == 'x' || mat[3] == 'o') {
		if mat[3] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	if mat[6] == mat[7] && mat[7] == mat[8] && (mat[6] == 'x' || mat[6] == 'o') {
		if mat[6] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}

	//verticales
	if mat[0] == mat[3] && mat[3] == mat[6] && (mat[0] == 'x' || mat[0] == 'o') {
		if mat[0] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	if mat[1] == mat[4] && mat[4] == mat[7] && (mat[1] == 'x' || mat[1] == 'o') {
		if mat[1] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	if mat[2] == mat[5] && mat[5] == mat[8] && (mat[2] == 'x' || mat[2] == 'o') {
		if mat[2] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}

	//diagonales
	if mat[0] == mat[4] && mat[4] == mat[8] && (mat[0] == 'x' || mat[0] == 'o') {
		if mat[0] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	if mat[2] == mat[4] && mat[4] == mat[6] && (mat[2] == 'x' || mat[2] == 'o') {
		if mat[0] == 'x' {
			return "Cliente"
		}
		return "Servidor"
	}
	return "z"
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	host := "localhost:8001"
	mat = []rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	ln, err = net.Listen("tcp", host)
	handlError(err)
	for {
		dibujarMatriz(mat)
		fmt.Println("Concexion establecida ....")

		conn, err = ln.Accept()
		handlError(err)

		r, _ := bufio.NewReader(conn).ReadString('\n')
		num, _ := strconv.Atoi(string(r[0]))
		mat[num] = 'x'
		if cantEspacios(mat) <= 1 {
			res := validarSiHayGanador(mat)
			if res != "z" {
				fmt.Println("!Ganador: " + res)
				enviar("!Ganador: " + res)
				break
			}
			fmt.Println("!Ganador: No hay")
			enviar("!Ganador: No hay")
			break
		}
		var ne bool
		for ne == false {
			r := rand.Intn(9)
			if mat[r] == 32 {
				mat[r] = 'o'
				ne = true
			}
		}

		res := validarSiHayGanador(mat)
		if res != "z" {
			dibujarMatriz(mat)
			fmt.Println("!Ganador: " + res)
			enviar("!Ganador: " + res)
			break
		}
		enviar(string(mat))
	}
	ln.Close()
	conn.Close()
}
