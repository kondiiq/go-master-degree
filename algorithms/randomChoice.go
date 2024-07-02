package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func RandomChoiceValue(maxCapacity int, knapsack []model.KnapsackItem) model.FinalKnapsack {
	var method string = "Random choice method"
	if maxCapacity == 0 || handler.IsArrayLenNull(knapsack) {
		return model.FinalKnapsack{
			Method: method,
			Value:  0,
		}
	}
	index, err := handler.ChooseRandomIndex(knapsack)
	if err != nil {
		handler.ErrorHandler(err)
	}
	currentResult := 0
	currentWeight := 0
	for currentWeight+knapsack[index].Weight <= maxCapacity || len(knapsack) == 1 {
		currentResult += knapsack[index].Value
		currentWeight += knapsack[index].Weight
		handler.RemoveElementWithIndexN(knapsack, index)
		index, err = handler.ChooseRandomIndex(knapsack)
		if err != nil {
			handler.ErrorHandler(err)
		}
	}
	return model.FinalKnapsack{
		Method: method,
		Value:  currentResult,
	}
}
