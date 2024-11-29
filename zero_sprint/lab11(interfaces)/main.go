package main

import (
	"fmt"
)

type Logger interface {
	Log(msg string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(msg string) {
	switch l.Level {
	case Error:
		fmt.Printf("ERROR: %s\n", msg)
	case Info:
		fmt.Printf("INFO: %s\n", msg)
	default:
		fmt.Printf("привет")
	}
}

//type Shape interface {
//	Area() float64
//}
//
//type Animal interface {
//	MakeSound()
//}
//
//type Dog struct {
//}
//
//func (d Dog) MakeSound() {
//	fmt.Println("Гав!")
//}
//
//type Cat struct {
//}
//
//func (c Cat) MakeSound() {
//	fmt.Println("Мяу!")
//}
//
//type Rectangle struct {
//	width  float64
//	height float64
//}
//
//func (r Rectangle) Area() float64 {
//	return r.width * r.height
//}
//
//type Circle struct {
//	radius float64
//}
//
//func (r Circle) Area() float64 {
//	return math.Pow(r.radius, 2) * math.Pi
//}

func main() {
	//figure1 := Circle{radius: 1.0}
	//fmt.Println(figure1.Area())
	//figure2 := Rectangle{width: 57.2, height: 10.2}
	//fmt.Println(figure2.Area())
	//
	//cat1 := Cat{}
	//cat1.MakeSound()
	//dog1 := Dog{}
	//dog1.MakeSound()

	errorLog := &Log{Level: Info}
	errorLog.Log("This is an error message")
	fmt.Println(Error)
}
