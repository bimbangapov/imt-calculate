package main

import "fmt"

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

	switch {
	case imt < 16:
		fmt.Printf("У вас недовес\n")
	case imt >= 16 && imt < 24.5:
		fmt.Printf("У вас нормальная масса тела\n")
	case imt >= 24.5 && imt < 30:
		fmt.Printf("У вас избыточный вес\n")
	case imt >= 30 && imt < 35:
		fmt.Printf("У вас ожирение первой степени\n")
	case imt >= 35 && imt < 40:
		fmt.Printf("У вас ожирение второй степени\n")
	case imt >= 40:
		fmt.Printf("У вас ожирение третьей степени\n")
	default:
		fmt.Printf("Неожиданное поведение системы\n")
	}

}
