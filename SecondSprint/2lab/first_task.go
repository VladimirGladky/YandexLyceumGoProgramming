package main

func Send(ch chan int, num int) {
	ch <- num
}
