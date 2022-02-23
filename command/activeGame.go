package command

import (
	"fmt"
	"joint-games/database"
	"joint-games/model"
	"joint-games/riot"
	"os"
)

func ActiveGame(summonerName string) {

	summoner := riot.GetSummoner(summonerName)
	if summoner.Id == "" {
		fmt.Println("The summoner was not found")
		os.Exit(2)
	}

	match, err := riot.GetActiveMatch(summoner.Id)

	if err != nil {
		fmt.Println("There are no active matches right now")
		os.Exit(2)
	}

	var resultMatches []model.Match

	for _, matchSummoner := range match.Summoners {

		matches, err := database.JointMatches(summoner, matchSummoner)
		if err != nil {
			fmt.Println(err)
			return
		}

		copy(resultMatches, matches)
	}

	for _, val := range resultMatches {
		fmt.Println(val.Id, val.Start)
	}
}
