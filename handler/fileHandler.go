package handler

import (
	"encoding/csv"
	"fmt"
	"knapsackProblem/model"
	"os"
	"strconv"
)

func Write2File(fileName string, content string) error{
	file , err := os.Create(fileName + ".txt")
	
	if err != nil {
		return err
	}
	defer file.Close()
	len, err := file.WriteString(content)
	
	if err != nil {
		return err
	}
	fmt.Printf("\nFile Name: %s", file.Name())
    fmt.Printf("\nLength: %d bytes", len)
	return nil
}

func ReadFromFile(sPath string) {
	fileName := sPath
	content, err := os.ReadFile(sPath)
	ErrorHandler(err)
	fmt.Printf("\nFile Name: %s", fileName)
    fmt.Printf("\nSize: %d bytes", len(content))
    fmt.Printf("\nData: %s", content)
}

func ReadCSVfile(sPath string) ([][]string, error) {
	file, err := os.Open(sPath)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, err
	}
	headers := records[0]
	fmt.Println("Headers:", headers)
	data := records[1:]
	return data, nil
}

func convert2Int(data [][]string) ([][]int, error){
	intData := make([][]int, len(data))
    for i, row := range data {
        intRow := make([]int, len(row))
        for j, value := range row {
            intValue, err := strconv.Atoi(value)
            if err != nil {
                return nil, err
            }
            intRow[j] = intValue
        }
        intData[i] = intRow
    }
    return intData, nil
}

func Convert2Knapsack(sPath string) ([]model.KnapsackItem, error) {
	var knapsack []model.KnapsackItem
	data, err := ReadCSVfile(sPath)
	if err != nil {
		return nil, err
	}
	result, err := convert2Int(data)
	if err!= nil {
        return nil, err
    }
	for _, row := range result {
        knapsack = append(knapsack, model.KnapsackItem{Weight: row[0], Value: row[1]})
    }
	return knapsack, nil
}

func Write2CSV(fileName string, data []model.FinalKnapsack) error {
	csvFile, err := os.Create(fileName + ".csv")
    
	if err != nil {
		fmt.Println("Something went wrong :", err)
		return err
	}
    defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
    defer writer.Flush()
	headers := []string{"Method", "Result"}
	
	if err := writer.Write(headers); err != nil {
		return err
	}
	
	for _, item := range data {
		record := []string{item.Method, strconv.Itoa(item.Value)}
		if err := writer.Write(record); err!= nil {
            fmt.Println(err)
			return err
        }
	}
	return nil
}

func Convert2JSON(data []model.FinalKnapsack) {}


