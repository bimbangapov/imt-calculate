package logic

import "fmt"

func ImtLogic(imt float64) string {
	switch {
	case imt < 16:
		return fmt.Sprintf("У вас недовес\n")
	case imt >= 16 && imt < 24.5:
		return fmt.Sprintf("У вас нормальная масса тела\n")
	case imt >= 24.5 && imt < 30:
		return fmt.Sprintf("У вас избыточный вес\n")
	case imt >= 30 && imt < 35:
		return fmt.Sprintf("У вас ожирение первой степени\n")
	case imt >= 35 && imt < 40:
		return fmt.Sprintf("У вас ожирение второй степени\n")
	case imt >= 40:
		return fmt.Sprintf("У вас ожирение третьей степени\n")
	default:
		return fmt.Sprintf("Неожиданное поведение системы\n")
	}
}
