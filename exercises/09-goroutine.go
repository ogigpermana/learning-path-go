package main

import (
	"fmt"
	"time"
)

func cetakPesan(nama string, ch chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- fmt.Sprintf("%s: pesan %d", nama, i+1)
	}
}

func main() {
	ch := make(chan string)
	
	go cetakPesan("Goroutine 1", ch)
	go cetakPesan("Goroutine 2", ch)
	
	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}
}