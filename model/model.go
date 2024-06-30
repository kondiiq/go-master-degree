package model

type KnapsackItem struct {
	Value int
	Weight int
}

type Result struct {
	Chromosome []bool
	Fitness int
}

type FinalKnapsack struct {
	Method string
	Value int
}
