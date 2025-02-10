package main

func Receive(ch chan int) int {
	return <-ch
}
