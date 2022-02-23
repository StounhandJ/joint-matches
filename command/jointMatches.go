package command

import (
	"fmt"
	"joint-games/database"
	"joint-games/riot"
	"os"
)

func JointMatches(summonerOneName string, summonerTwoName string) {

	summonerOne := riot.GetSummoner(summonerOneName)
	if summonerOne.Id == "" {
		fmt.Println("The summonerOne was not found")
		os.Exit(2)
	}

	summonerTwo := riot.GetSummoner(summonerTwoName)
	if summonerTwo.Id == "" {
		fmt.Println("The summonerTwo was not found")
		os.Exit(2)
	}

	matches, err := database.JointMatches(summonerOne, summonerTwo)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range matches {
		fmt.Println(val)
	}
}
