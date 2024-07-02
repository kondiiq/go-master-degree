package handler

import (
	"errors"
	"fmt"
	"io"
	"knapsackProblem/model"
	"log"
	"math/rand"
	"os"
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
	if IsArrayLenNull(arr) {
		return knapsack, 0, errors.New("entered values and weights are not of same length")
	}
	if len(arr) == 1 {
		return arr[0], 0, nil
	}
	max := arr[0]
	index := 0
	for item := 1; item < len(arr); item++ {
		max, index = findMaxAndIndex(arr, max, index, item, property)
	}
	return max, index, nil
}

func AppendItemsIntoKnapsack(values []int, weights []int) ([]model.KnapsackItem, error) {
	var finalKnapsack []model.KnapsackItem
	if len(values) != len(weights) {
		return finalKnapsack, errors.New("entered values and weights are not of same length")
	}
	for item := 0; item < len(values); item++ {
		item := model.KnapsackItem{
			Value:  values[item],
			Weight: weights[item],
		}
		finalKnapsack = append(finalKnapsack, item)
	}
	return finalKnapsack, nil
}

func findMaxAndIndex(arr []model.KnapsackItem, max model.KnapsackItem, index int, currIndex int, property string) (model.KnapsackItem, int) {
	if property == "Weight" || property == "weight" {
		if arr[currIndex].Weight > max.Weight {
			max = arr[currIndex]
			index = currIndex
			return max, index
		}
	}
	if property == "Value" || property == "value" {
		if arr[currIndex].Value > max.Value {
			max = arr[currIndex]
			index = currIndex
			return max, index
		}
	}
	return max, index
}

func FindMaxValueWeightRatioItemAndIndex(arr []float32) int {
	if IsArrayLenNull(arr) {
		return 0
	}
	var index int
	max := arr[0]
	for item := 0; item < len(arr); item++ {
		if arr[item] > max {
			max = arr[item]
			index = item
		}
	}
	return index
}

func ChooseRandomIndex(knapsack []model.KnapsackItem) (int, error) {
	if IsArrayLenNull(knapsack) {
		return 0, errors.New("current list has not data")
	}
	index := rand.Intn(len(knapsack))
	for knapsack[index].Value == 0 && knapsack[index].Weight == 0 {
		index = rand.Intn(len(knapsack))
	}
	return index, nil
}

func IsArrayLenNull[T comparable](arr []T) bool {
	return reflect.ValueOf(arr).IsZero()
}

func IsCapacityNull[T comparable](capacity T) bool {
	return reflect.ValueOf(capacity).IsZero()
}

func CreateValueWeightRatio(knapsack []model.KnapsackItem) []float32 {
	var result []float32
	var partResult float32
	for item := 0; item < len(knapsack); item++ {
		if reflect.ValueOf(float32(knapsack[item].Value)).IsZero() && reflect.ValueOf(float32(knapsack[item].Weight)).IsZero() {
			partResult = 0
		} else {
			partResult = float32(knapsack[item].Value) / float32(knapsack[item].Weight)
		}
		result = append(result, partResult)
	}
	return result
}

func CompareValues(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CalculateWeightAndValueOfKnapsack(result model.Result, knapsack []model.KnapsackItem) (int, int) {
	if len(knapsack) == 0 {
		return 0, 0
	}
	currentWeight := 0
	currentValue := 0
	for item := range knapsack {
		if !result.Chromosome[item] {
			currentWeight += 0
			currentValue += 0
		} else {
			currentValue += knapsack[item].Value
			currentWeight += knapsack[item].Weight
		}
	}
	return currentWeight, currentValue
}

func CalculateFitness(knapsack []model.KnapsackItem, result model.Result, maxCapacity, penaltyValue int) int {
	knapsackWeight, knapsackValue := CalculateWeightAndValueOfKnapsack(result, knapsack)
	if knapsackWeight > maxCapacity {
		return knapsackValue - penaltyValue
	}
	return knapsackValue
}

func ReadFile(path string) {
	file, err := os.Open(path)
	ErrorHandler(err)

	defer file.Close()
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buffer[:]))
		}
	}
}
