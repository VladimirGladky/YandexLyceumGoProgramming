package main

import (
	"fmt"
)

func main() {
	inputJSON := []byte(`[
        {
            "name": "Oleg",
            "class": "9B"
        },
        {
            "name": "Ivan",
            "class": "9A"
        },
        {
            "name": "Maria",
            "class": "9B"
        },
        {
            "name": "John",
            "class": "9A"
        }
    ]`)

	classJSONMap, err := splitJSONByClass(inputJSON)
	if err != nil {
		fmt.Println("Пидор")
	} else {
		print(classJSONMap)
	}
}
