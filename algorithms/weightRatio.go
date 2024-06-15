package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func WeightRatioValue(maxCapacity int, knapsack []model.KnapsackItem) int {
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack) {
		return 0
	}
	currentWeight := 0
	currentResult := 0
	ratioArray := handler.CreateValueWeightRatio(knapsack)
	index := handler.FindMaxValueWeightRatioItemAndIndex(ratioArray)
	for currentWeight +  knapsack[index].Weight <= maxCapacity || len(knapsack) == 1 {
		currentResult += knapsack[index].Value
		currentWeight += knapsack[index].Weight
		handler.RemoveElementWithIndexN(knapsack, index)
		ratioArray = handler.CreateValueWeightRatio(knapsack)
		index = handler.FindMaxValueWeightRatioItemAndIndex(ratioArray)
	}
	return currentResult
}