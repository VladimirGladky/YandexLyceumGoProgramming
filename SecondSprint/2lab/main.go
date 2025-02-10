package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		Send(ch, 13)
	}()
	val := Receive(ch)
	fmt.Println(val)
}
