package raindrops

import (
	"strconv"
	"strings"
)

// Convert Takes in a int value returns a string value depending on the int value given.
func Convert(i int) string {
	word := ""
	// Word is a empty string ready to get appended values if is divisible by the int

	if i%3 == 0 {
		word += "Pling"
	}
	if i%5 == 0 {
		word += "Plang"
	}
	if i%7 == 0 {
		word += "Plong"
	}
	/*ContainsAny Function checks value inside word to check if it matches
	  if true returns word value.
	*/
	if strings.ContainsAny("PlingPlangPlong", word) {
		return word
	}

	return strconv.Itoa(i)
}
