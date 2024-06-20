package algorithms

import (

	"knapsackProblem/model"
	rand "math/rand"
	"sort"
)

func GAResolver(noGeneration, maxCapacity int, mutationRate float32, knapsack []model.KnapsackItem) int {
	Population := initSinglePopulation(len(knapsack), noGeneration)
	for i := range Population{
		Population[i].Fitness = calculateFitness(knapsack, Population[i], maxCapacity)
		Population  = selectAndReproduce(Population)
	}
	_, bestResultValue := calculateWeightAndValueOfKnapsack(getBestIndividual(Population), knapsack)
	return bestResultValue
}

func initSinglePopulation(populationSize, noPopulation int) []model.Result {
    population := make([]model.Result, noPopulation)

    for element := range noPopulation {
        chromosome := make([]bool, populationSize)
        for elementProperty := range chromosome {
            chromosome[elementProperty] = rand.Intn(2) == 1
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

func rouletteSelection(population []model.Result, cumulativeFitness []int, totalFitness int) model.Result {
	pickedItem := rand.Intn(totalFitness)
	idx := sort.Search(len(cumulativeFitness), func(item int) bool {
		return cumulativeFitness[item] >= pickedItem
	})
	return population[idx]
}

func selectAndReproduce(population []model.Result) []model.Result {
	newPopulation := make([]model.Result, len(population))
	cumulativeFitness := make([]int, len(population))
	cumulativeFitness[0] = population[0].Fitness
	for i := 1; i < len(population); i++ {
		cumulativeFitness[i] = cumulativeFitness[i - 1] + population[i].Fitness
	}
	totalFitness := cumulativeFitness[len(cumulativeFitness) - 1]

	for i := range newPopulation {
		parent1 := rouletteSelection(population, cumulativeFitness, totalFitness)
        parent2 := rouletteSelection(population, cumulativeFitness, totalFitness)
        newPopulation[i] = crossover(parent1, parent2)
	}
	return newPopulation
}

func crossover(parent1, parent2 model.Result) model.Result {
    crossoverPoint := rand.Intn(len(parent1.Chromosome))
    childChromosome := append(parent1.Chromosome[:crossoverPoint], parent2.Chromosome[crossoverPoint:]...)
    return model.Result{Chromosome: childChromosome}
}

func getBestIndividual(population []model.Result) model.Result {
    best := population[0]
    for _, individual := range population {
        if individual.Fitness > best.Fitness {
            best = individual
        }
    }
    return best
}