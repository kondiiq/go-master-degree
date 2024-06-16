package model

type KnapsackItem struct {
	Value int
	Weight int
}

type Resolution struct {
	Chromosome []bool
	Fitness int
}
