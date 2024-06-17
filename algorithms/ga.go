package algorithms

import (
	"fmt"
	"knapsackProblem/model"
	r "math/rand"
)

func GAResolver(noGeneration, maxCapacity int, mutationRate float32, knapsack []model.KnapsackItem) string {
	Population := initSinglePopulation(len(knapsack), noGeneration)
	for i := range Population{
		Population[i].Fitness = calculateFitness(knapsack, Population[i], maxCapacity)
	}
	sortedPopulation := sortByFitness(Population)
	
	return ""
}

func initSinglePopulation(populationSize, noPopulation int) []model.Result {
    population := make([]model.Result, noPopulation)

    for element := range noPopulation {
        chromosome := make([]bool, populationSize)
        for elementProperty := range chromosome {
            chromosome[elementProperty] = r.Intn(2) == 1
        }
        population[element] = model.Result{Chromosome: chromosome}
    }
    return population
}

func calculateWeightAndValueOfKnapsack(result model.Result, knapsack []model.KnapsackItem) (int, int) {
	if len(knapsack) == 0 {
		return 0, 0
	}
	currentWeight := 0
	currentValue  := 0
	for i := range knapsack {
		if !result.Chromosome[i] {
			currentWeight += 0
			currentValue  += 0
		} else{
			currentValue += knapsack[i].Value
			currentWeight += knapsack[i].Weight
		}
	}
	return currentWeight, currentValue
}

func calculateFitness(knapsack []model.KnapsackItem, result model.Result, maxCapacity int) int {
	knapsackWeight, knapsackValue := calculateWeightAndValueOfKnapsack(result, knapsack)
	if knapsackWeight > maxCapacity {
        return knapsackValue - 1000
    }
	return knapsackValue
}

func sortByFitness(Population []model.Result) []model.Result {
	for i := range Population {
        for j := range Population {
            if Population[i].Fitness > Population[j].Fitness {
                Population[i], Population[j] = Population[j], Population[i]
            }
        }
    }
    return Population
}