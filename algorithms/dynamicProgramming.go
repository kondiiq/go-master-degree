package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func DynamicProgramming(maxCapacity int, knapsack []model.KnapsackItem) int {
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack){ 
		return 0
	}
	maxValues := make([]int, maxCapacity + 1)
	for _, item := range knapsack {
		for c := maxCapacity; c >= item.Weight; c-- {
			// Aktualizacja tablicy maxValues
			maxValues[c] = handler.CompareValues(maxValues[c], maxValues[c-item.Weight] + item.Value)
		}
	}
	return maxValues[maxCapacity]
}