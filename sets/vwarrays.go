package sets

import (
	"knapsackProblem/handler"
	"knapsackProblem/model"
)

func CreateKnapsack15() ([]model.KnapsackItem, error) {
	values := Value15()
	weights := Weight15()
	var knapsackItems []model.KnapsackItem
	for i := 0; i < len(values); i++ {
		knapsackItems = append(knapsackItems, model.KnapsackItem{Value: values[i], Weight: weights[i]})
	}
	return knapsackItems, nil
}

func CreateKnapsack50() ([]model.KnapsackItem, error) {
	knapsack, err := handler.Convert2Knapsack("./data/dataset50.csv")
	if err != nil {
		return nil, err
	}
	return knapsack, nil
}

func CreateKnapsack100() ([]model.KnapsackItem, error) {
	knapsack, err := handler.Convert2Knapsack("./data/dataset100.csv")
	if err != nil {
		return nil, err
	}
	return knapsack, nil
}

func CreateKnapsack200() ([]model.KnapsackItem, error) {
	knapsack, err := handler.Convert2Knapsack("./data/dataset200.csv")
	if err != nil {
		return nil, err
	}
	return knapsack, nil
}

func Value15() []int {
	return []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
}

func Weight15() []int {
	return []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
}

func MaxCapacity15() int {
	return 750
}

func Penalty15() int {
	return 1500
}

func Temp15() float32 {
	return 1000.0
}

func MaxCapacity50() int {
	return 75
}

func Penalty50() int {
	return 1500
}

func Temp50() float32 {
	return 1000.0
}
