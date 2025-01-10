package main

import (
	"fmt"
	"math"
	"time"
)

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {
	return len(t.notes)
}

func (t ToDoList) CountTopPrioritiesTasks() int {
	var sum int
	for i := range t.tasks {
		if t.tasks[i].IsTopPriority() {
			sum++
		}
	}
	return sum
}

func (t ToDoList) CountOverdueTasks() int {
	var sum int
	for i := range t.tasks {
		if t.tasks[i].IsOverdue() {
			sum++
		}
	}
	return sum
}

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s", p.name, p.age, p.address)
}

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {
	return time.Now().After(t.deadline)
}

func (t Task) IsTopPriority() bool {
	return t.priority == 4 || t.priority == 5
}

func (e Employee) CalculateTotalSalary() {
	totalSalary := math.Round((e.salary+e.bonus)*100) / 100
	fmt.Printf("Employee: %s\nPosition: %s\nTotal Salary: %.2f", e.name, e.position, totalSalary)
}

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	var sum float64
	for i := 0; i < s.solvedProblems; i++ {
		sum += s.scoreForOneTask
	}
	return sum > s.passingScore
}

func main() {
	//man := Person{name: "Vova", age: 12, address: "Rostov"}
	//man.Print()

	man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
	man.CalculateTotalSalary()
}
