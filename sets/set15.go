package sets

import (
	"fmt"
	"knapsackProblem/algorithms"
	"knapsackProblem/handler"
)

func Set15Items(weightSlice, valueSlice []int, maxCapacity, penaltyValue int, initTemperature float32) {
	
	noGenerations := 100
	knapsack, err := handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	noItems := len(knapsack)
	rResult := algorithms.RandomChoiceValue(maxCapacity, knapsack)
	fmt.Println("Random alg result is :", rResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	bfResult := algorithms.BruteForce(maxCapacity, noItems, knapsack)
	fmt.Println("Brute force result is :", bfResult, "$")
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	mvResult := algorithms.MaxValue(maxCapacity, knapsack)
	fmt.Println("max valueSlice result is :", mvResult)
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	mwResult := algorithms.MaxWeight(maxCapacity, knapsack)
	fmt.Println("max weightSlice result is :", mwResult)
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	wgResult := algorithms.WeightRatioValue(maxCapacity, knapsack)
	fmt.Println("Weight ratio result is :", wgResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	dpResult := algorithms.DynamicProgramming(maxCapacity, knapsack)
	fmt.Println("Dynamic programming result is :", dpResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	gaResult := algorithms.GAResolver(noGenerations, maxCapacity, penaltyValue, 0.01, knapsack)
	fmt.Println("GA result is :", gaResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(valueSlice[:], weightSlice[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	saResult := algorithms.SimulatedAnnealing(knapsack, maxCapacity, noGenerations, penaltyValue, 0.01, initTemperature)
	fmt.Println("SA result is :", saResult.Fitness)
}