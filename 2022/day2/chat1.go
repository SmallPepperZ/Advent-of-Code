package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// A map mapping each possible hand shape to a score.
	// Rock has a score of 1, Paper has a score of 2, and Scissors has a score of 3.
	handScore := map[rune]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}

	// A map mapping each possible outcome of a round to a score.
	// A loss is 0 points, a draw is 3 points, and a win is 6 points.
	outcomeScore := map[string]int{
		"loss": 0,
		"draw": 3,
		"win":  6,
	}

	// The total score is initially 0.
	totalScore := 0

	// Read the strategy guide from standard input.
	// Each line of the input is a separate round, and each line contains
	// two characters separated by a space: the first character is what
	// the opponent will play, and the second character is what you should
	// play.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Split the line into the opponent's hand shape and your hand shape.
		var opponentShape, yourShape rune
		fmt.Sscanf(scanner.Text(), "%c %c", &opponentShape, &yourShape)

		// Calculate the score for the hand shape you chose.
		shapeScore := handScore[yourShape]

		// Calculate the score for the outcome of the round.
		// If your hand shape beats the opponent's hand shape, the round is a win.
		// If your hand shape is the same as the opponent's hand shape, the round is a draw.
		// Otherwise, the round is a loss.
		var outcome string
		switch opponentShape {
		case 'A':
			if yourShape == 'Z' {
				outcome = "win"
			} else if yourShape == 'X' {
				outcome = "draw"
			} else {
				outcome = "loss"
			}
		case 'B':
			if yourShape == 'X' {
				outcome = "win"
			} else if yourShape == 'Y' {
				outcome = "draw"
			} else {
				outcome = "loss"
			}
		case 'C':
			if yourShape == 'Y' {
				outcome = "win"
			} else if yourShape == 'Z' {
				outcome = "draw"
			} else {
				outcome = "loss"
			}
		}
		outcomeScore := outcomeScore[outcome]

		// Add the scores for the round to the total score.
		totalScore += shapeScore + outcomeScore
	}

	// Print the total score.
	fmt.Println(totalScore)
}
