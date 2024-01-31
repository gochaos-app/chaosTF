package opsChaos

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func Test() {
	choices := []struct {
		value  string
		weight int
	}{
		{"foo", 3},
		{"bar", 1},
		{"baz", 2},
	}

	// Prompt the user for input
	var userInput int
	fmt.Print("Enter a number: ")
	fmt.Scan(&userInput)

	// Find the choice with the corresponding weight
	var matchingChoices []struct {
		value  string
		weight int
	}
	for _, choice := range choices {
		if choice.weight == userInput {
			matchingChoices = append(matchingChoices, choice)
			fmt.Println(matchingChoices)
		}
	}

	// If there are matching choices, select one randomly
	if len(matchingChoices) > 0 {
		// Calculate the total weight of all matching choices
		totalWeight := 0
		for _, choice := range matchingChoices {
			totalWeight += choice.weight
		}

		// Generate a random number between 0 and the total weight
		rand.Seed(42) // Seed the random number generator for reproducibility
		randomNumber := rand.Intn(totalWeight)

		// Select the matching choice based on the random number and its weight
		weightSum := 0
		for _, choice := range matchingChoices {
			weightSum += choice.weight
			if randomNumber < weightSum {
				fmt.Println(choice.value)
				break
			}
		}
	} else {
		// If there are no matching choices, select any choice randomly
		// Calculate the total weight of all choices
		totalWeight := 0
		for _, choice := range choices {
			totalWeight += choice.weight
		}

		// Generate a random number between 0 and the total weight
		rand.Seed(42) // Seed the random number generator for reproducibility
		randomNumber := rand.Intn(totalWeight)

		// Select the choice based on the random number and its weight
		weightSum := 0
		for _, choice := range choices {
			weightSum += choice.weight
			if randomNumber < weightSum {
				fmt.Println(choice.value)
				break
			}
		}
	}
}
