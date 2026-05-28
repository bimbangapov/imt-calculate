package main

import (
	"fmt"
	"imt-calculate/internal/logic"
)

func main() {
	fmt.Println("Программа для рассчета ИМТ - индекса массы тела")
	var inputWeight, inputHeight, imt float64

	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&inputWeight)

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&inputHeight)
	inputHeight = inputHeight / 100

	fmt.Printf("Ваш рост: %.2f м.\n"+
		"Ваш вес: %.2f кг. \n", inputHeight, inputWeight)

	imt = inputWeight / (inputHeight * inputHeight)
	fmt.Printf("Ваш ИМТ = %.2f \n", imt)

	fmt.Println(logic.ImtLogic(imt))
}
