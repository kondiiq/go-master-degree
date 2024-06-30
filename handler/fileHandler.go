package handler

import (
	"encoding/csv"
	"fmt"
	"os"
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

func ReadCSVfile(sPath string) {
	file, err := os.Open(sPath)
	ErrorHandler(err)
	reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
    fmt.Println(records)
}

func Convert2Struct() {
	// Convert CSV to Struct
}

func Write2CSV(fileName string, data [][]string) error {
	// Write content to csv
	csvFile, err := os.Create(fileName + ".csv")
    if err != nil {
		fmt.Errorf("Something went wrong :", err)
		return err
	}
    defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	if err := csvWriter.WriteAll(data); err != nil {
        return fmt.Errorf("could not write data to CSV: %v", err)
    }
	csvWriter.Flush()
    return nil
}


