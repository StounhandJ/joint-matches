package command

import (
	"fmt"
	"joint-games/database"
	"joint-games/model"
	"joint-games/riot"
	"os"
)

func ActiveGame(summonerName string) {

	db := database.NewDataBase()
	db.Db.Model(&model.MatchSummoner{})

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
		var matches []model.Match
		_, err := db.Db.Query(
			&matches,
			"select matches.id, matches.match_id, matches.start "+
				"from match_summoners as main "+
				"join match_summoners as dop on main.match_id = dop.match_id "+
				"join matches on matches.id = main.match_id "+
				"where dop.summoner_id = ? and main.summoner_id = ?", matchSummoner.Id, summoner.Id)
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
