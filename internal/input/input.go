package input

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func InputWeightHeight() (float64, float64, error) {
	var weight, height float64
	var inputWeight, inputHeight string

	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&inputWeight)
	inputWeight = strings.ReplaceAll(inputWeight, ",", ".")
	weight, err := strconv.ParseFloat(inputWeight, 64)
	if err != nil {
		return 0, 0, err
	}
	if weight <= 0 {
		return 0, 0, errors.New("Рост/вес должны быть положительными")
	}

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&inputHeight)
	inputHeight = strings.ReplaceAll(inputHeight, ",", ".")
	height, err = strconv.ParseFloat(inputHeight, 64)
	if err != nil {
		return 0, 0, err
	}

	height = height / 100
	if height <= 0 {
		return 0, 0, errors.New("Рост/вес должны быть положительными")
	}

	return height, weight, nil
}
