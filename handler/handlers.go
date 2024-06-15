package handler 

import (
	"log"
	"knapsackProblem/model"
	"errors"
	"math/rand"
	"reflect"
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
	knapsack[index].Value = 0
	knapsack[index].Weight = 0
	return knapsack, nil
}

func FindMaxPropertyItem(arr []model.KnapsackItem, property string) (model.KnapsackItem, int, error) {
	knapsack := model.KnapsackItem{}
	if(IsArrayLenNull(arr)){
		return knapsack, 0, errors.New("entered values and weights are not of same length")
	}
	if(len(arr) == 1){
		return arr[0], 0, nil
	}
	max := arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		max, index = findMaxAndIndex(arr, max, index, i, property)
	}
	return max, index, nil
}

func AppendItemsIntoKnapsack(values []int, weights []int) ([]model.KnapsackItem, error) {
	var finalKnapsack  []model.KnapsackItem
	if len(values) != len(weights) {
		return finalKnapsack, errors.New("entered values and weights are not of same length")
	}
	for i := 0; i < len(values); i++ {
		item := model.KnapsackItem{
			Value: values[i],
			Weight: weights[i],
		}
        finalKnapsack = append(finalKnapsack, item)
    }
	return finalKnapsack, nil
}

func  findMaxAndIndex(arr []model.KnapsackItem, max model.KnapsackItem, index int, currIndex int, property string) (model.KnapsackItem, int) {
	if property == "Weight" || property == "weight" {
		if arr [currIndex].Weight > max.Weight {
			max = arr[currIndex]
			index = currIndex
			return max, index
		}
	}
	if property == "Value" || property == "value" {
		if arr [currIndex].Value > max.Value {
			max = arr[currIndex]
			index = currIndex
			return max, index
		}
	}
	return max, index
}

func FindMaxValueWeightRatioItemAndIndex(arr []float32) int{
	if IsArrayLenNullf32(arr) {
		return 0
	}
	var index int
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
            max = arr[i]
			index = i
        }
	}
	return index
}

func ChooseRandomIndex(knapsack []model.KnapsackItem) (int, error) {
	if IsArrayLenNull(knapsack) {
        return 0, errors.New("current list has not data")
    }
	index := rand.Intn(len(knapsack))
	for knapsack[index].Value == 0 &&  knapsack[index].Weight == 0 {
		index = rand.Intn(len(knapsack))
	}
    return index, nil
}

func IsArrayLenNull(arr []model.KnapsackItem) bool {
	return len(arr) == 0
}

func IsArrayLenNullf32(arr []float32) bool {
	return len(arr) == 0
}

func CreateValueWeightRatio(knapsack []model.KnapsackItem) []float32 {
	var result []float32
	var partResult float32
	for i := 0; i < len(knapsack); i++ {
		if reflect.ValueOf(float32(knapsack[i].Value)).IsZero() &&  reflect.ValueOf(float32(knapsack[i].Weight)).IsZero() {
			partResult = 0
		} else {
			partResult = float32(knapsack[i].Value) /  float32(knapsack[i].Weight)
		}
		result = append(result, partResult)
	}
	return result
}