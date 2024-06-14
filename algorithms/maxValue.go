package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

// func MaxValue(maxCapacity int, noItems int, knapsack []model.KnapsackItem) int {
// 	if maxCapacity == 0 || noItems == 0 {
// 		return 0
// 	}
// 	currentWeight := 0
// 	currentResult := 0
// 	for currentWeight <= maxCapacity || len(knapsack) == 1 {
// 		mValue, index, err := handler.FindMaxPropertyItem(knapsack, "value")
// 		if err != nil {
// 			handler.ErrorHandler(err)
// 		}
// 		currentResult += mValue.Value
// 		currentWeight += mValue.Weight
// 		handler.RemoveElementWithIndexN(knapsack, index)
// 	}
// 	return currentResult
// }

func MaxValue(maxCapacity int, noItems int, knapsack []model.KnapsackItem) int {
	if maxCapacity == 0 || noItems == 0 {
		return 0
	}
	currentWeight := 0
	currentResult := knapsack[noItems - 1].Value
	currentWeight += knapsack[noItems - 1].Weight
	noItems -= 1
	handler.RemoveElementWithIndexN(knapsack, noItems - 1)
	for currentWeight +  knapsack[noItems - 1].Weight <= maxCapacity || len(knapsack) == 1 {
		currentResult += knapsack[noItems - 1].Value
		currentWeight += knapsack[noItems - 1].Value
		handler.RemoveElementWithIndexN(knapsack, noItems - 1)
		noItems -= 1
	}
	return currentResult
}