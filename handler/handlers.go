package handler 

import (
	"log"
	"knapsackProblem/model"
	"errors"
)

func ErrorHandler(err error) {
	log.Fatal("Something went wrong: \n", err)
	panic(err)
}

func RemoveElementWithIndexN(knapsack []model.KnapsackItem, index int) ([]model.KnapsackItem, error) {
	if index >= len(knapsack) {
		log.Fatal("Index out of range")
		return nil, errors.New("entered values and weights are not of same length")
	}
	return append(knapsack[:index], knapsack[index + 1:]...), nil
}