package main

import (
	"fmt"
	"knapsackProblem/algorithms"
	"knapsackProblem/handler"
)

func main() {
	value := []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
	weight := []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
	max_capaxity := 750
	knapsack, err := handler.AppendItemsIntoKnapsack(value[:], weight[:])
	if err != nil {
		handler.ErrorHandler(err)
	}
	noItems := len(knapsack)
	fmt.Println("Knapsack is :", knapsack, " with limit :", max_capaxity, "kg and ", len(knapsack), "  items")
	for i := 0; i < len(knapsack); i++ {
		fmt.Println("Item:", i + 1," in knapsack have value : ", knapsack[i].Value, "$ and weight:", knapsack[i].Weight, "kg")
	}
	bfResult := algorithms.BruteForce(max_capaxity, noItems, knapsack)
	fmt.Println("Brute force result is :", bfResult, "$")
	mvResult := algorithms.MaxValue(max_capaxity, knapsack)
	fmt.Println("max value choose result is :", mvResult)
}

