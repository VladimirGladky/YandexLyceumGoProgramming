package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	var w World
	if height > 0 && width > 0 {
		cells := make([][]bool, height)
		for i := range cells {
			cells[i] = make([]bool, width)
		}
		w = World{height, width, cells}
	} else {
		return nil, errors.New("неправильный ввод")
	}

	return &w, nil
}

func run() error {
	m := os.Args[1]
	n := os.Args[2]
	p := os.Args[3]
	width, err := strconv.Atoi(m)
	if err != nil || width <= 0 {
		return fmt.Errorf("ширина должна быть положительным числом")
	}

	height, err := strconv.Atoi(n)
	if err != nil || height <= 0 {
		return fmt.Errorf("высота должна быть положительным числом")
	}

	fillPercentage, err := strconv.Atoi(p)
	if err != nil || fillPercentage < 0 || fillPercentage > 100 {
		return fmt.Errorf("процент заполнения должен быть числом от 0 до 100")
	}
	f, err := os.Create("config.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	configLine := fmt.Sprintf("%dx%d %d%%", width, height, fillPercentage)
	_, err = f.WriteString(configLine)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	if err := run(); err != nil {
		println("Error:", err.Error())
	}
}
