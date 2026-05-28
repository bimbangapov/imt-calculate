package input

import (
	"errors"
	"fmt"
)

func InputWeightHeight() (float64, float64, error) {
	var weight, height float64

	fmt.Print("Введите свой вес в килограммах: ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		return 0, 0, errors.New("Вес должен быть числом")
	}

	if weight <= 0 {
		return 0, 0, errors.New("Рост/вес должны быть положительными")
	}

	fmt.Print("Введите свой рост в сантиметрах: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		return 0, 0, errors.New("Рост должен быть числом")
	}

	height = height / 100
	if height <= 0 {
		return 0, 0, errors.New("Рост/вес должны быть положительными")
	}

	return height, weight, nil
}
