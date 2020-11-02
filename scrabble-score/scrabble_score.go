package scrabble

import "strings"

//Score is a method that takes in word as a string and returns the value as an Int
func Score(word string) int {
	//word is now all uppercase
	word = strings.ToUpper(word)

	tiles := make(map[string]int)
	tiles["AEIOULNRST"] = 1
	tiles["DG"] = 2
	tiles["BCMP"] = 3
	tiles["FHVWY"] = 4
	tiles["K"] = 5
	tiles["JX"] = 8
	tiles["QZ"] = 10

	var score int
	// Range gets the length of the word by ilteration
	for _, x := range word {

		//checking the Key and Values String and Int
		for k, v := range tiles {

			// If statement checks the Key against Rune value created via Range for loop
			// converted back to a string so its checking String for String
			if strings.Contains(k, string(x)) {
				// score is then added to by the value of K
				score += v
			}
		}
		// Then goes to the next character in word given if not returns score
	}
	// Returns the updated score value
	return score
}
