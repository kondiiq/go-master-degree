package main

import (
	"knapsackProblem/sets"
)

func main() {
	value := []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
	weight := []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
	var maxCapaxity int = 750
	var penaltyValue int = 1000
	var initTemperature float32 = 1000.0
	sets.Set15Items(weight, value, maxCapaxity, penaltyValue, initTemperature)
}



