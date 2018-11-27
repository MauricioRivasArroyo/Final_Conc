package main

import (
    "fmt"
    "time"
)

var end chan bool

func philosopher(id int, fi, fi1 chan bool) {
    for {
        fmt.Printf("Filósofo %d pensando.\n", id)
        time.Sleep(time.Millisecond*500)
        <-fi
        <-fi1
        fmt.Printf("Filósofo %d comiendo.\n", id)
        time.Sleep(time.Millisecond*500)
        fi<- true
        fi1<- true
    }
    end<- true
}
func fork(fi chan bool) {
    for {
        fi<- true
        time.Sleep(time.Millisecond*500)
        <-fi
    }
}

func main() {
    n := 5
    forks := make([]chan bool, n)
    for i := 0; i < n; i++ {
        forks[i] = make(chan bool)
    }
    for i := 0; i < n; i++ {
        go philosopher(i, forks[i], forks[(i + 1) % n])
        go fork(forks[i])
    }
    for i := 0; i < n; i++ {
        <-end
    }
}
