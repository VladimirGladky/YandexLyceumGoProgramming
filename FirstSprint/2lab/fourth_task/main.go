package main

import "fmt"

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Speed float64
	Type  string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	res := make(map[string]float64)
	for _, vehicle := range vehicles {
		var Type string
		switch v := vehicle.(type) {
		case Car:
			Type = v.Type
		case Motorcycle:
			Type = v.Type
		default:
			continue
		}
		if Type == "Мотоцикл" {
			res["main.Motorcycle"] = vehicle.CalculateTravelTime(distance)
		} else {
			res["main.Car"] = vehicle.CalculateTravelTime(distance)
		}
	}
	return res
}

func main() {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)
	fmt.Println(travelTimes)
}
