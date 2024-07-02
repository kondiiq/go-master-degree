package algorithms

import (
	"knapsackProblem/model"
)

func BruteForce(maxCapacity int, noItems int, knapsack []model.KnapsackItem) int {
	if noItems == 0 || maxCapacity == 0 {
		return 0
	}
	if knapsack[noItems-1].Weight > maxCapacity {
		return BruteForce(maxCapacity, noItems-1, knapsack)
	}
	return max(BruteForce(maxCapacity, noItems-1, knapsack), knapsack[noItems-1].Value+BruteForce(maxCapacity-knapsack[noItems-1].Weight, noItems-1, knapsack))
}

func ConvertBFtoFinalKnapsack(bfResult int) model.FinalKnapsack {
	return model.FinalKnapsack{
		Method: "Brute force method",
		Value:  bfResult,
	}
}
