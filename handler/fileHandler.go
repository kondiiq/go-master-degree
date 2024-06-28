package handler

import(
	"fmt"
    "os"
	"encoding/csv"
)

func Write2File(fileName string, content string){
	file , err := os.Create(fileName + ".txt")
	ErrorHandler(err)

	defer file.Close()

	len, err := file.WriteString(content)

	ErrorHandler(err)

	fmt.Printf("\nFile Name: %s", file.Name())
    fmt.Printf("\nLength: %d bytes", len)
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

func Write2CSV(fileName string, data []byte, delimitter string) {
	// Write content to csv
}


