package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func DynamicProgramming(maxCapacity int, knapsack []model.KnapsackItem) model.FinalKnapsack {
	var method string = "Dynamic programming method"
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack) {
		return model.FinalKnapsack{
			Method: method,
			Value:  0,
		}
	}
	maxValues := make([]int, maxCapacity+1)
	for _, item := range knapsack {
		for c := maxCapacity; c >= item.Weight; c-- {
			// Aktualizacja tablicy maxValues
			maxValues[c] = handler.CompareValues(maxValues[c], maxValues[c-item.Weight]+item.Value)
		}
	}
	return model.FinalKnapsack{
		Method: method,
		Value:  maxValues[maxCapacity],
	}
}
