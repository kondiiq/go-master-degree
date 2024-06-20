package algorithms

import (

	"knapsackProblem/model"
	"knapsackProblem/handler"
	rand "math/rand"
	"sort"
)

func GAResolver(noGeneration, maxCapacity, penaltyValue int, mutationRate float32, knapsack []model.KnapsackItem) int {
	Population := initSinglePopulation(len(knapsack), noGeneration)
	for i := range Population{
		Population[i].Fitness = handler.CalculateFitness(knapsack, Population[i], maxCapacity, 1000)
		Population  = selectAndReproduce(Population)
	}
	_, bestResultValue := handler.CalculateWeightAndValueOfKnapsack(getBestIndividual(Population), knapsack)
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