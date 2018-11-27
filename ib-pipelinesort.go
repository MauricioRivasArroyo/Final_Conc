package main

import (
    "fmt"
    "math"
)

func p(in, out, sorted chan float64) {
    nros := make([]float64, 0, 10)
    menor := math.Inf(1)
    pos := -1
    i := 0
    for nro := range in {
        if nro < menor {
            menor = nro
            pos = i
        }
        i++
        nros = append(nros, nro)
    }
    for i, nro := range nros {
        if i != pos {
            out<- nro
        }
    }
    close(out)
    sorted<- menor
}
func main() {
    c1 := make(chan float64)
    c2 := make(chan float64)
    c3 := make(chan float64)
    c4 := make(chan float64)
    c5 := make(chan float64)
    cs := make(chan float64)

    go p(c1, c2, cs)
    go p(c2, c3, cs)
    go p(c3, c4, cs)
    go p(c4, c5, cs)

    c1<- 6
    c1<- 1
    c1<- 5
    c1<- 4
    close(c1)
    fmt.Println(<-cs)
    fmt.Println(<-cs)
    fmt.Println(<-cs)
    fmt.Println(<-cs)


}
