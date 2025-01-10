package main

import (
	"fmt"
	"time"
)

func main() {
	//ModifyFile("text1.txt", 5, "body")
	log, err := ExtractLog("text1.txt", time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return
	}
	fmt.Println(log)
}
