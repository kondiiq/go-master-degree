package main

import (
	"fmt"
	"knapsackProblem/algorithms"
	"knapsackProblem/handler"
	"knapsackProblem/model"
	"errors"
)

func main() {
	value := []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
	weight := []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
	max_capaxity := 750
	knapsack, err := appendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	fmt.Println("Knapsack is :", knapsack, " with limit :", max_capaxity, "kg and ", len(knapsack), "  items")
	for i := 0; i < len(knapsack); i++ {
		fmt.Println("Item:", i + 1," in knapsack have value : ", knapsack[i].Value, "$ and weight:", knapsack[i].Weight, "kg")
	}
	noItems := len(value)
	bfResult := algorithms.BruteForce(max_capaxity, noItems, knapsack)
	fmt.Println("Brute force result is :", bfResult, "$")
	mvResult := algorithms.MaxValue(max_capaxity, noItems, knapsack)
	fmt.Println("max value choose result is :", mvResult)
}

func appendItemsIntoKnapsack(values []int, weights []int) ([]model.KnapsackItem, error) {
	var finalKnapsack  []model.KnapsackItem
	if len(values) != len(weights) {
		return finalKnapsack, errors.New("entered values and weights are not of same length")
	}
	for i := 0; i < len(values); i++ {
		item := model.KnapsackItem{
			Value: values[i],
			Weight: weights[i],
		}
        finalKnapsack = append(finalKnapsack, item)
    }
	return finalKnapsack, nil
}