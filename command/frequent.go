package command

import (
	"fmt"
	"joint-games/database"
	"joint-games/model"
	"joint-games/riot"
	"os"
	"strconv"
)

func Frequent() {
	db := database.NewDataBase()

	if len(os.Args) == 2 {
		fmt.Println("Specify the player's nickname as the first parameter")
		os.Exit(2)
	}

	summoner := riot.GetSummoner(os.Args[2])
	if summoner.Id == "" {
		fmt.Println("The summoner was not found")
		os.Exit(2)
	}

	countGame := 3
	if len(os.Args) >= 4 {
		countGame, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("The start indent must be a number")
			os.Exit(2)
		}
	}

	var g []model.FrequentSummoner

	_, err := db.Db.Query(
		&g,
		"select summ.id, summ.puuid, summ.name, COUNT(*) from match_summoners as main "+
			"join match_summoners as dop on main.match_id = dop.match_id and main.summoner_id != ? "+
			"join summoners as summ on summ.id = main.summoner_id "+
			"where dop.summoner_id = ? "+
			"GROUP BY summ.id HAVING COUNT(*) >= ? "+
			"ORDER BY \"count\" DESC;", summoner.Id, summoner.Id, countGame)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range g {
		fmt.Println(val.Name, val.Count)
	}

}
