package main

import (
	"fmt"
)

func findStudentByID(id int, students map[int]string) (string, error) {
	name, found := students[id]
	if !found {
		return "", fmt.Errorf("студент с ID %d не найден", id)
	}

	return name, nil
}

func FindMaxKey(m map[int]int) int {
	res := -1000000000
	for key, _ := range m {
		if key > res {
			res = key
		}
	}
	return res
}

func SumOfValuesInMap(m map[int]int) int {
	var res int
	for _, value := range m {
		res += value
	}
	return res
}

func SwapKeysAndValues(m map[string]string) map[string]string {
	res := make(map[string]string)
	for key, value := range m {
		res[value] = key
	}
	return res
}

func CountingSort(contacts []string) map[string]int {
	res := make(map[string]int)
	for _, num := range contacts {
		res[num]++
	}
	return res
}

func DeleteLongKeys(m map[string]int) map[string]int {
	res := make(map[string]int)
	for key, value := range m {
		if len(key) >= 6 {
			res[key] = value
		}
	}
	return res
}

func main() {
	m := map[string]string{
		"Яндекс":        "+74957397000",
		"Музей Яндекса": "+74991101133",
	}
	fmt.Println(SwapKeysAndValues(m))
}
