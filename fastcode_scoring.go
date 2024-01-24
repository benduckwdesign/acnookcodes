package acnookcodes

import (
	"unicode"
)

// TODO: Sophisticate for symbols, number of finger movements, letter distance on keyboard, etc.
func scoreSequentialCharacters(code string) int {
	maxSequential := 0
	currentSequential := 0
	bonusFourteen := false

	charTypeBonus := 0

	for i := 0; i < len(code)-1; i++ {
		if unicode.IsLetter(rune(code[i])) {
			charTypeBonus += 2
			if unicode.IsDigit(rune(code[i+1])) {
				charTypeBonus -= 1
			} else if unicode.IsLetter(rune(code[i+1])) {

			} else {
				charTypeBonus -= 2
			}
		} else if unicode.IsDigit(rune(code[i])) {
			charTypeBonus += 1
			if unicode.IsDigit(rune(code[i+1])) {

			} else if unicode.IsLetter(rune(code[i+1])) {
				charTypeBonus -= 1
			} else {
				charTypeBonus -= 2
			}
		} else {
			charTypeBonus += 0
			if unicode.IsDigit(rune(code[i+1])) {
				charTypeBonus -= 1
			} else if unicode.IsLetter(rune(code[i+1])) {
				charTypeBonus -= 1
			} else {

			}
		}

		if code[i] == code[i+1] {
			currentSequential++
			if currentSequential > 13 {
				bonusFourteen = true
			}
			if currentSequential > maxSequential {
				maxSequential = currentSequential
			}
		} else {
			currentSequential = 0
		}
	}

	// Adjust the scoring criteria as per your requirements.
	// You can return a higher score for more sequential characters.
	score := maxSequential
	score += charTypeBonus
	if bonusFourteen {
		score += 2
	}
	return score
}
