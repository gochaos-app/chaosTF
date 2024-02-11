package opsChaos

import (
	"time"

	"golang.org/x/exp/rand"
)

type Action struct {
	Name   string
	Weight int
}

func ExecuteAction(userInput string) string {
	actions := []Action{
		{Name: "stop", Weight: 1},
		{Name: "reboot", Weight: 1},
		{Name: "terminate", Weight: 3},
		{Name: "shutdown", Weight: 2},
	}
	if userInput == "" {
		userInput = "kill"
	}
	var totalWeight int
	for _, action := range actions {
		if userInput == "basic" && (action.Name == "stop" || action.Name == "reboot") {
			totalWeight += action.Weight
		} else if userInput == "kill" && (action.Name == "terminate" || action.Name == "shutdown") {
			totalWeight += action.Weight
		}
	}
	RandomSeed := rand.NewSource(uint64(time.Now().UnixMicro()))
	rand.Seed(RandomSeed.Uint64())
	randomNumber := rand.Intn(totalWeight)

	var cumulativeWeight int
	for _, action := range actions {
		cumulativeWeight += action.Weight
		if randomNumber < cumulativeWeight {

			return action.Name

		}
	}
	return ""

}
