package sets

import (
	"knapsackProblem/algorithms"
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func SolveKnapsack(weightSlice, valueSlice []int, maxCapacity, penaltyValue, noGenerations int, initTemperature float32) ([]model.FinalKnapsack, error) {
	var resultToWrite []model.FinalKnapsack

	knapsack, err := handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	noItems := len(knapsack)

	rResult := algorithms.RandomChoiceValue(maxCapacity, knapsack)
	resultToWrite = append(resultToWrite, rResult)

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	bfResult := algorithms.ConvertBFtoFinalKnapsack(algorithms.BruteForce(maxCapacity, noItems, knapsack))
	resultToWrite = append(resultToWrite, bfResult)

	if err != nil {
		handler.ErrorHandler(err)
	}

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	mvResult := algorithms.MaxValue(maxCapacity, knapsack)
	resultToWrite = append(resultToWrite, mvResult)

	if err != nil {
		handler.ErrorHandler(err)
	}

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	mwResult := algorithms.MaxWeight(maxCapacity, knapsack)
	resultToWrite = append(resultToWrite, mwResult)

	if err != nil {
		handler.ErrorHandler(err)
	}

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	wgResult := algorithms.WeightRatioValue(maxCapacity, knapsack)
	resultToWrite = append(resultToWrite, wgResult)

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	dpResult := algorithms.DynamicProgramming(maxCapacity, knapsack)
	resultToWrite = append(resultToWrite, dpResult)

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	gaResult := algorithms.GAResolver(noGenerations, maxCapacity, penaltyValue, 0.01, knapsack)
	resultToWrite = append(resultToWrite, gaResult)

	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])

	if err != nil {
		handler.ErrorHandler(err)
	}
	saResult := algorithms.Convert2FinalKnapsack(algorithms.SimulatedAnnealing(knapsack, maxCapacity, noGenerations, penaltyValue, 0.01, initTemperature).Fitness, "Simulated annealling(SA) mehod")
	resultToWrite = append(resultToWrite, saResult)

	return resultToWrite, nil
}
