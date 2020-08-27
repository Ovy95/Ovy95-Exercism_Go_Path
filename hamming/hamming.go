package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("length of strings isn't equal")
	}
	StrandsA := []rune(a)
	Strandsb := []rune(b)

	counter := 0

	for i := range StrandsA {
		if StrandsA[i] != Strandsb[i] {
			counter++
		}
	}
	return counter, nil
}
