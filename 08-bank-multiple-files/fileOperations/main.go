package fileOperations

import (
	"errors"
	"fmt"
	"os"
)

func WriteFloatToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	valueText, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("value file not found")
	}
	var value float64
	_, parsingError := fmt.Sscan(string(valueText), &value)
	if parsingError != nil {
		return 0, errors.New("value file is invalid")
	}
	return value, nil
}
