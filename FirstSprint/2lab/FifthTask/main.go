package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

func CreateReport(user User, reportDate string) Report {
	id := user.ID
	if user.ID <= 0 {
		id = rand.Intn(1000)
	}
	return Report{id, reportDate, user}
}

func PrintReport(report Report) {
	fmt.Printf("Пользователь:%s Email:%s Возраст:%v ", report.User.Name, report.User.Email, report.User.Age)
	fmt.Printf("Дата создания отчёта: %s", report.Date)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	res := make([]Report, 0, len(users))
	for _, user := range users {
		res = append(res, CreateReport(user, reportDate))
	}

	return res
}

func main() {
	user := User{ID: 1, Name: "Иван", Email: "ivan@example.com", Age: 30}
	reportDate := time.Now().Format("2006-01-02")

	report := CreateReport(user, reportDate)
	fmt.Println(report)
}
