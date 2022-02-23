package command

import (
	"fmt"
	"joint-matches/database"
	"joint-matches/riot"
	"os"
)

func Frequent(summonerName string, countGame int) {

	summoner := riot.GetSummoner(summonerName)
	if summoner.Id == "" {
		fmt.Println("The summoner was not found")
		os.Exit(2)
	}

	frequentSummoners, err := database.FrequentSummonerInMatch(summoner, countGame)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range frequentSummoners {
		fmt.Println(val.Name, val.Count)
	}

}
