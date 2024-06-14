package algorithms

import (
	"fmt"
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

func WeightRatioValue(maxCapacity int, knapsack []model.KnapsackItem) int {
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack) {
		return 0
	}
	noItems := len(knapsack)
	currentWeight := 0
	currentResult := knapsack[noItems - 1].Value
	currentWeight += knapsack[noItems - 1].Weight
	noItems -= 1
	valueWeightRatioArray := createValueWeightRatio(knapsack)
	index := handler.FindMaxValueWeightRatioItemAndIndex(valueWeightRatioArray)
	handler.RemoveElementWithIndexN(knapsack, noItems)
	fmt.Print(index)
	for currentWeight +  knapsack[noItems - 1].Weight <= maxCapacity || len(knapsack) == 1 {
		currentResult += knapsack[noItems - 1].Value
		currentWeight += knapsack[noItems - 1].Weight
		handler.RemoveElementWithIndexN(knapsack, noItems)
		noItems -= 1
	}
	return currentResult
}

func createValueWeightRatio(knapsack []model.KnapsackItem) []float32 {
	var result []float32
	for i := 0; i < len(knapsack); i++ {
		item := float32(knapsack[i].Value) /  float32(knapsack[i].Weight)
		result = append(result, item)
	}
	return result
}