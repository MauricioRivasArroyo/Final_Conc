package main

import (
	"fmt"
)

var ch chan int

func sort(array []int) {
	for v := range ch {
		for i := 0; i < len(array); i++ {
			if array[v] < array[i] {
				aux := array[v]
				array[v] = array[i]
				array[i] = aux
			}
		}
	}
}

func main() {
	array := []int{3, 9, 1, 2, 8}
	ch = make(chan int)

	go sort(array)
	for i, v := range array {
		fmt.Printf("valor: %d\n", v)
		ch <- i
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
