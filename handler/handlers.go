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
	ret := make([]model.KnapsackItem, 0)
	ret = append(ret, knapsack[:index]...)
	return append(ret, knapsack[index + 1:]...), nil
}

func FindMaxPropertyItem(arr []model.KnapsackItem, property string) (model.KnapsackItem, int, error) {
	knapsack := model.KnapsackItem{}
	if(len(arr) == 0){
		return knapsack, 0, errors.New("entered values and weights are not of same length")
	}
	if(len(arr) == 1){
		return arr[0], 0, nil
	}
	max := arr[0]
	index := 0
	if(property == "Value" || property == "value"){
		for i := 1; i < len(arr); i++ {
			if arr [i].Value > max.Value {
				max = arr[i]
				index = i
			}
		}
		return max, index, nil
	}
	if(property == "Weight" || property == "weight"){
		for i := 1; i < len(arr); i++ {
			if arr [i].Weight > max.Weight {
				max = arr[i]
				index = i
			}
		}
		return max, index, nil
	}
	return max, index, nil
}