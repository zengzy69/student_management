package util

import (
	"encoding/csv"
	"log"
	"os"
)

// ReadCSV 读取CSV文件
func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading CSV:", err)
		return nil, err
	}

	return records, nil
}
