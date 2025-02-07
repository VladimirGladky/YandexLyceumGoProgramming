package main

import "slices"

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val) // Количество драгоценностей

	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}
	numberMatrix := make([][][]int, n+1)
	for i := range numberMatrix {
		numberMatrix[i] = make([][]int, maxWeight+1)
		for j := range numberMatrix[i] {
			numberMatrix[i][j] = []int{}
		}
	}

	for item := 1; item <= n; item++ { // Переберём все предметы из сундука
		for capacity := 1; capacity <= maxWeight; capacity++ {
			// Всё ниже — о рюкзаке вместимостью capacity
			var numbersWithCurrent []int
			maxcostWithoutCurrent := matrix[item-1][capacity] //  Максимальная стоимость предыдущих предметов
			maxcostWithCurrent := 0                           //  Для хранения максимальной стоимости, если положим текущий предмет

			weightOfCurrent := chest.wt[item-1] // Масса текущего
			if capacity >= weightOfCurrent {    // Проверяем, влезет ли текущий предмет в рюкзак
				// Если текущий влез, смотрим, что ещё взять
				maxcostWithCurrent = chest.val[item-1] // Сначала положим текущий предмет
				numbersWithCurrent = append(numbersWithCurrent, item-1)
				remainingCapacity := capacity - weightOfCurrent         // Проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // Максимальная стоимость оставшегося места
				numbersWithCurrent = append(numbersWithCurrent, numberMatrix[item-1][remainingCapacity]...)
			}

			if maxcostWithoutCurrent > maxcostWithCurrent {
				matrix[item][capacity] = maxcostWithoutCurrent
				numberMatrix[item][capacity] = numberMatrix[item-1][capacity]
			} else {
				matrix[item][capacity] = maxcostWithCurrent
				numberMatrix[item][capacity] = numbersWithCurrent
			}
		}
	}
	slices.Sort(numberMatrix[n][maxWeight])
	return matrix[n][maxWeight], numberMatrix[n][maxWeight]
}
