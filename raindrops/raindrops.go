package raindrops

import (
	"strconv"
)

func Convert(i int) string {

	// Created remainder as feel its easier to read, Bit unnecessary, feel it reads easier
	remainder := 0
	word := ""

	// Word is a empty string ready to get appended values if is divisible by the int
	if i%3 == remainder {
		word += "Pling"
	}
	if i%5 == remainder {
		word += "Plang"
	}
	if i%7 == remainder {
		word += "Plong"
	}
	if i%3 != remainder && i%5 != remainder && i%7 != remainder {
		return strconv.Itoa(i)
		// Uses strcov package to convert int to string returns just value
	}
	return word

}
