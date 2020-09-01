package raindrops

import "strconv"

func Convert(i int) string {

	// Created remainder as feel its easier to read, Bit unnecessary, feel it reads easier
	remainder := 0
	// checks each that int that gets passed is divisible by more
	//then one value before singler to get the right result

	if i%3 == remainder && i%5 == remainder && i%7 == remainder {
		return "PlingPlangPlong"
	} else if i%5 == remainder && i%7 == remainder {
		return "PlangPlong"
	} else if i%3 == remainder && i%7 == remainder {
		return "PlingPlong"
	} else if i%3 == remainder && i%5 == remainder {
		return "PlingPlang"
	} else if i%3 == remainder {
		return "Pling"
	} else if i%5 == remainder {
		return "Plang"
	} else if i%7 == remainder {
		return "Plong"
		// else uses strcov package to convert int to string
	} else {
		return strconv.Itoa(i)
	}

}
