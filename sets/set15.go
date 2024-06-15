package sets

import (
	"fmt"
	"knapsackProblem/algorithms"
	"knapsackProblem/handler"
)

func Set15Items() {
	value := []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
	weight := []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
	max_capaxity := 750
	knapsack, err := handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	noItems := len(knapsack)
	rResult := algorithms.RandomChoiceValue(max_capaxity, knapsack)
	fmt.Println("Random alg result is :", rResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	bfResult := algorithms.BruteForce(max_capaxity, noItems, knapsack)
	fmt.Println("Brute force result is :", bfResult, "$")
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	mvResult := algorithms.MaxValue(max_capaxity, knapsack)
	fmt.Println("max value result is :", mvResult)
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	mwResult := algorithms.MaxWeight(max_capaxity, knapsack)
	fmt.Println("max weight result is :", mwResult)
	if err != nil {
		handler.ErrorHandler(err)
	}
	knapsack, err = handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	wgResult := algorithms.WeightRatioValue(max_capaxity, knapsack)
	fmt.Println("Weight ratio result is :", wgResult)
	knapsack, err = handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	dpResult := algorithms.DynamicProgramming(max_capaxity, knapsack)
	fmt.Println("Dynamic programming result is :", dpResult)
}