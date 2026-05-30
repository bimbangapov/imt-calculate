package input

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func InputWeightHeight() (float64, float64, error) {
	var inputWeight, inputHeight string

	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&inputWeight)

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&inputHeight)

	return ParseWeightHeight(inputWeight, inputHeight)
}

func ParseWeightHeight(inputWeight, inputHeight string) (float64, float64, error) {
	inputWeight = strings.ReplaceAll(inputWeight, ",", ".")
	weight, err := strconv.ParseFloat(inputWeight, 64)
	if err != nil {
		return 0, 0, err
	}
	if weight <= 0 {
		return 0, 0, errors.New("Рост/Вес должен быть положительным")
	}

	inputHeight = strings.ReplaceAll(inputHeight, ",", ".")
	height, err := strconv.ParseFloat(inputHeight, 64)
	if err != nil {
		return 0, 0, err
	}
	if height <= 0 {
		return 0, 0, errors.New("Рост/Вес должен быть положительным")
	}

	return height, weight, nil
}
