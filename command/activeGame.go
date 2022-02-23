package command

import (
	"fmt"
	"joint-matches/database"
	"joint-matches/model"
	"joint-matches/riot"
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
		if summoner.Id == matchSummoner.Id {
			continue
		}

		matches, err := database.JointMatches(summoner, matchSummoner)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := range matches {
			matches[i].Summoners = []model.Summoner{matchSummoner}
		}

		resultMatches = append(resultMatches, matches...)
	}

	for _, val := range resultMatches {
		fmt.Println(val.Summoners[0], " | ", val)
	}
}
