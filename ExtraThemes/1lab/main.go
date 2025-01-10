package main

import (
	"errors"
	"fmt"
	"slices"
)

// 3
var (
	positionOrder = map[string]int{
		"директор":       1,
		"зам. директора": 2,
		"начальник цеха": 3,
		"мастер":         4,
		"рабочий":        5,
	}
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	workers []Worker
}

type Worker struct {
	Name            string
	Position        string
	Salary          uint
	ExperienceYears uint
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if name == "" || position == "" {
		return errors.New("неправильный ввод")
	}
	c.workers = append(c.workers, Worker{Name: name, Position: position, Salary: salary, ExperienceYears: experience})
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	var res []string
	slices.SortFunc(c.workers, func(a, b Worker) int {
		if a.Salary*a.ExperienceYears > b.Salary*b.ExperienceYears {
			return -1
		} else if a.Salary*a.ExperienceYears < b.Salary*b.ExperienceYears {
			return 1
		} else {
			if positionOrder[a.Position] < positionOrder[b.Position] {
				return -1
			} else if positionOrder[a.Position] > positionOrder[b.Position] {
				return 1
			}
		}
		return 0
	})

	for _, x := range c.workers {
		res = append(res, fmt.Sprintf("%v - %v - %v", x.Name, x.Salary*x.ExperienceYears, x.Position))
	}
	return res, nil
}

// 1
func SortNums(nums []uint) {
	slices.Sort(nums)
}

// 2
func SortNames(names []string) {
	slices.Sort(names)
}

// 4
func SortAndMerge(left, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)
	var res []int
	var leftPoint int
	var rightPoint int
	for leftPoint < len(left) && rightPoint < len(right) {
		if left[leftPoint] < right[rightPoint] {
			res = append(res, left[leftPoint])
			leftPoint++
		} else {
			res = append(res, right[rightPoint])
			rightPoint++
		}
	}
	if leftPoint == len(left) {
		res = append(res, right[rightPoint:]...)
	} else {
		res = append(res, left[leftPoint:]...)
	}
	return res
}

// 5
func SortByFreq(str string) string {
	Freq := make(map[rune]int)
	var res string
	for _, x := range str {
		if _, exists := Freq[x]; exists {
			Freq[x]++
		} else {
			Freq[x] = 1
		}
	}

	type kv struct {
		Key   rune
		Value int
	}

	var sortedPairs []kv
	for k, v := range Freq {
		sortedPairs = append(sortedPairs, kv{k, v})
	}

	slices.SortFunc(sortedPairs, func(a, b kv) int {
		if a.Value < b.Value {
			return -1
		} else if a.Value > b.Value {
			return 1
		} else {
			if a.Key < b.Key {
				return -1
			} else {
				return 1
			}
		}
	})

	for _, x := range sortedPairs {
		for i := 0; i < x.Value; i++ {
			res += string(x.Key)
		}
	}

	return res
}

func main() {
	//nums := []uint{7, 5, 10, 4, 1, 2, 1, 2}
	//names := []string{"fnjkdkfd", "fdkjnfjkn", "aaa"}
	//SortNames(names)
	//fmt.Println(names)
	//SortNums(nums)
	//fmt.Println(nums)

	fmt.Println(SortAndMerge([]int{2, 4, 1}, []int{2, 1, 3, 6, 5}))

	fmt.Println(SortByFreq("abbbzzzatt"))
}
