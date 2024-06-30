package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func MaxValue(maxCapacity int, knapsack []model.KnapsackItem) model.FinalKnapsack {
	var method string = "Max value items mathod"
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack) {
		return  model.FinalKnapsack{
			Method : method,
			Value : 0,
		}
	}
	noItems := len(knapsack)
	currentResult := knapsack[noItems - 1].Value
	currentWeight := knapsack[noItems - 1].Weight
	handler.RemoveElementWithIndexN(knapsack, noItems  - 1)
	_, index, err := handler.FindMaxPropertyItem(knapsack, "value")
	if err!= nil {
        handler.ErrorHandler(err)
    }
	for currentWeight +  knapsack[index].Weight <= maxCapacity || len(knapsack) == 1 {
		currentResult += knapsack[index].Value
		currentWeight += knapsack[index].Weight
		handler.RemoveElementWithIndexN(knapsack, index)
		_, index, err = handler.FindMaxPropertyItem(knapsack, "value")
	}
	return model.FinalKnapsack{
		Method : method,
		Value : currentResult,
	}
}