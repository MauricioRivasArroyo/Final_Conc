package main

import (
    "fmt"
)

const MAX = 9
const K = 4

func compress(inC, pipe chan rune) {
    n := 0
    previous := <-inC
    for c := range inC {
        if c == previous && n < MAX - 1 {
            n++
        } else {
            if n > 0 {
                pipe<- rune(n+1 + 48)
                n = 0
            }
            pipe<- previous
            previous = c
        }
    }
    pipe<- previous
    close(pipe)
}

func output(pipe, outC chan rune) {
    m := 0
    for c := range pipe {
        outC<- c
        m++
        if m >= K {
            outC<- '\n'
            m = 0
        }
    }
    close(outC)
}

func main() {
    inC := make(chan rune)
    pipe := make(chan rune)
    outC := make(chan rune)
    go compress(inC, pipe)
    go output(pipe, outC)
    go func() {
        msg := "ABBBCDEEEEFFGGGGGHIJJJKKLLLSL"
        for _, c := range msg {
            inC<- c
        }
        close(inC)
    }()
    for c := range outC {
        fmt.Printf("%c", c)
    }
    fmt.Println()
}
