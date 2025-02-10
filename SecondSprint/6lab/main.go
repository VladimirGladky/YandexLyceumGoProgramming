package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(TimeoutFibonacci(5, time.Second*3))
}
