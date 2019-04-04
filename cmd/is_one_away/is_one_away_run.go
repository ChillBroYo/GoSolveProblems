package main

/* One away edits can be grouped into 3 categories: add a character, remove a character and
 * replace a character. Write a method that checks if a string is one or zero edits
 * away from another
 */

import (
	"log"
	"math"
)

func main() {
	log.Println(isAtMostOneAway("Hello", "Hello"))
	log.Println(isAtMostOneAway("hello", "Hello"))
	log.Println(isAtMostOneAway("noway", "Hello"))
	log.Println(isAtMostOneAway("helloo", "Hello"))
}

// First Impression solution
func isAtMostOneAway(referenceString string, stringToCheck string) bool {
	rCheckMap := map[byte]int{}
	nCheckMap := map[byte]int{}
	isOneAway := false

	for i := 0; i < len(referenceString); i++ {
		rCheckMap[referenceString[i]]++
	}

	for i := 0; i < len(stringToCheck); i++ {
		nCheckMap[stringToCheck[i]]++
	}

	for char, cardinality := range rCheckMap {
		if math.Abs(float64(nCheckMap[char]-cardinality)) > 1.0 {
			return false
		} else if math.Abs(float64(nCheckMap[char]-cardinality)) == 1.0 && isOneAway {
			return false
		} else if math.Abs(float64(nCheckMap[char]-cardinality)) == 1.0 && !isOneAway {
			isOneAway = true
		}
	}

	if len(nCheckMap) >= 2+len(rCheckMap) {
		return false
	} else if len(nCheckMap) == 1+len(rCheckMap) && isOneAway {
		return false
	}

	return true
}
