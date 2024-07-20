package test

import(
	"knapsackProblem/handler"
	"knapsackProblem/model"
	"knapsackProblem/sets"
	"testing"
	"fmt"
)

func TestCompareValues(t *testing.T) {
	testCases := []struct{
		a, b int
        expected int
	}{
		{5, 10, 10},
		{10, 5, 10},
		{10, 10, 10},
		{-1, -2, -1},
		{0, -1, 0},
		{-1, 1, 1},
	}
	
	for index, test := range testCases{
		t.Run(fmt.Sprintf("Test case CompareValues foo No.%d", index) ,func(t *testing.T) {
			result := handler.CompareValues(test.a, test.b)
			if result != test.expected {
				t.Errorf("CompareValues(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
			}
			fmt.Printf("Test TestCompareValues with no.%d fly with color :)", index)
		})
	}
}

func TestIsCapacityNull(t *testing.T) {
	testCases := []struct{
		itemWeight int
        expected bool
	}{
		{5, false},
		{0, true},
		{-1, false},
	}
	
	for index, test := range testCases{
		t.Run(fmt.Sprintf("Test case IsCapacityNull foo No.%d", index) ,func(t *testing.T) {
			result := handler.IsCapacityNull(test.itemWeight)
			if result != test.expected {
				t.Errorf("IsCapacityNull(%d) = %v; expected %v", test.itemWeight, result, test.expected)
			}
			fmt.Printf("Test TestIsCapacityNull with no.%d fly with color :)", index)
		})
	}
}

func TestCreateValueWeightRatio(t *testing.T) {
	values := sets.Value15()
	weights := sets.Weight15()
	results := make([]float32, 0 )
	results = append(results, 1.928571429, 1.904109589, 1.935064935, 1.875, 1.902439024, 1.873563218, 1.922222222, 1.957446809, 1.959183673, 1.896226415, 1.909090909, 1.89380531, 1.92173913, 1.940677966, 2.00)
	knapsack, _ := handler.AppendItemsIntoKnapsack(values, weights)  
	testCases := []struct{
		testKnapsack []model.KnapsackItem
        expected []float32
	}{
		{knapsack, results},
	}
	
	for index, test := range testCases{
		t.Run(fmt.Sprintf("Test case CreateValueWeightRatio foo No.%d", index) ,func(t *testing.T) {
			result := handler.CreateValueWeightRatio(test.testKnapsack)
			if result[index] != test.expected[index] {
				t.Error("CreateValueWeightRatio", test.testKnapsack, result, test.expected)
			}
			fmt.Printf("Test CreateValueWeightRatio with no.%d fly with color :)", index)
		})
	}
}
