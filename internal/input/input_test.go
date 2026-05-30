package input

import "testing"

func TestParseWeightHeightValidInput(t *testing.T) {
	height, weight, err := ParseWeightHeight("70", "175")

	if err != nil {
		t.Fatalf("Ожидали ошибку nil а получили: %v", err)
	}

	if height != 175 {
		t.Errorf("Ожидали рост 1.75, а получили: %.2f", height)
	}

	if weight != 70 {
		t.Errorf("Ожидали вес 70, а получили: %.2f", weight)
	}
}

func TestParseWeightHeightInvalidInput(t *testing.T) {
	_, _, err := ParseWeightHeight("abc", "175")
	if err == nil {
		t.Fatal("ожидали ошибку, получили nil")
	}
}

func TestParseWeightHeightInvalidInput2(t *testing.T) {
	_, _, err := ParseWeightHeight("-70", "175")
	if err == nil {
		t.Fatal("Ожидали ошибку отрицательного веса/роста получили nil ")
	}
}
