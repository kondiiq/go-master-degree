package algorithms

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
	"math"
	"math/rand"
)

func SimulatedAnnealing(knapsack []model.KnapsackItem, maxCapacity, maxIteration, penaltyValue int, coolingRate, initTemperature float32) model.Result {
	currentSolution := generateSolution(len(knapsack))
	currentSolution.Fitness = handler.CalculateFitness(knapsack, currentSolution, maxCapacity, penaltyValue)
	bestSolution := currentSolution
	temp := initTemperature

	for element := 0; element < maxIteration; element++ {
		neighbour := getNeigbour(currentSolution)
		neighbour.Fitness = handler.CalculateFitness(knapsack, neighbour, maxCapacity, penaltyValue)

		if neighbour.Fitness > currentSolution.Fitness || rand.Float32() < acceptancePobability(currentSolution.Fitness, neighbour.Fitness, temp) {
			currentSolution = neighbour
		}
		temp *= coolingRate
		if temp < 1e-3 {
			break
		}
	}
	return bestSolution
}

func Convert2FinalKnapsack(result int, method string) model.FinalKnapsack {
	return model.FinalKnapsack{
		Method: method,
		Value:  result,
	}
}

func generateSolution(knapsackSize int) model.Result {
	chromosome := make([]bool, knapsackSize)
	for element := range chromosome {
		chromosome[element] = rand.Intn(2) == 1
	}
	return model.Result{Chromosome: chromosome}
}

func getNeigbour(solution model.Result) model.Result {
	neighbour := model.Result{Chromosome: make([]bool, len(solution.Chromosome))}
	copy(neighbour.Chromosome, solution.Chromosome)
	idx := rand.Intn(len(solution.Chromosome))
	neighbour.Chromosome[idx] = !neighbour.Chromosome[idx]
	return neighbour
}

func acceptancePobability(currentFitness, newFitness int, temperature float32) float32 {
	if newFitness > currentFitness {
		return 1.0
	}
	return float32(math.Exp(float64(newFitness-currentFitness) / float64(temperature)))
}
