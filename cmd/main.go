package main

import (
	"fmt"
	"imt-calculate/internal/input"
	"imt-calculate/internal/logic"
)

func main() {
	fmt.Println("Программа для рассчета ИМТ - индекса массы тела")
	for {
		inputHeight, inputWeight, err := input.InputWeightHeight()
		if err != nil {
			fmt.Printf("Неккоректно введены рост/вес, ошибка: %v\n\n", err)
			continue
		}
		inputHeight = inputHeight / 100
		fmt.Printf("Ваш вес: %.2f кг.\n"+
			"Ваш рост: %.2f м. \n", inputWeight, inputHeight)

		imt := inputWeight / (inputHeight * inputHeight)
		fmt.Printf("Ваш ИМТ = %.2f \n", imt)

		fmt.Println(logic.ImtLogic(imt))
	}
}
