package main

import (
	"fmt"
	"knapsackProblem/handler"
	"knapsackProblem/sets"
	"knapsackProblem/db"
)

func main() {

	value := sets.Value15()
	weight := sets.Weight15()
	var maxCapaxity int = sets.MaxCapacity15()
	var penaltyValue int = sets.Penalty15()
	var initTemperature float32 = sets.Temp15()
	const noGenerations int = 10000
	knapsack, err := sets.SolveKnapsack(weight, value, maxCapaxity, penaltyValue, noGenerations, initTemperature)
	if err != nil {
		fmt.Print(err)
	}
	handler.Write2CSV("result", knapsack)

	handler.RunBackend()
	db.ConnectDatabase()
}