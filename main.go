package main

import (
	"knapsackProblem/sets"
)

func main() {
	value := sets.Value15()
	weight := sets.Weight15()
	var maxCapaxity int = sets.MaxCapacity15()
	var penaltyValue int = sets.Penalty15()
	var initTemperature float32 = sets.Temp15()
	const noGenerations int = 10000
	sets.SolveKnapsack(weight, value, maxCapaxity, penaltyValue, noGenerations, initTemperature)
}