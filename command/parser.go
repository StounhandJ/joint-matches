package command

import (
	"fmt"
	"joint-matches/database"
	"joint-matches/model"
	"joint-matches/riot"
	"os"
)

var (
	db database.DataBase
)

func Parser(summonerName string, start int, update bool) {

	summoner := riot.GetSummoner(summonerName)
	if summoner.Id == "" {
		fmt.Println("The summoner was not found")
		os.Exit(2)
	}

	db = *database.NewDataBase()
	db.Db.Model(&model.MatchSummoner{})

	var summoners []model.Summoner
	var startTime int64

	startTime = -1
	if update {
		lastMatch, err := database.LastMatch(summoner)
		if err != nil {
			fmt.Println("The last match is not found, the parsing of all begins")
		} else {
			startTime = lastMatch.Start.Unix()
		}
	}

	i := 0
	for match := range riot.GetMatches(summoner, start, startTime) {
		i += 1

		_ = db.Db.Model(&match).Where("match_id = ?", match.Id).Select()
		if match.IdDB != 0 {
			fmt.Println(fmt.Sprintf("%d Match(%s) already added", i, match.Id))
			continue
		}

		match.Summoners = getSummoners(&match.Summoners, &summoners)
		summoners = append(summoners, match.Summoners...)

		saveMatch(&match)

		fmt.Println(fmt.Sprintf("%d Match(%s) saved", i, match.Id))
	}
}

func getSummoners(matchSummoners *[]model.Summoner, summoners *[]model.Summoner) []model.Summoner {
	var summonersMatch []model.Summoner

	for _, summoner := range *matchSummoners {
		var s model.Summoner
		id := findSummoner(summoners, &summoner)
		if id == -1 {
			_ = db.Db.Model(&s).Where("puuid = ?", summoner.Puuid).Select()
			if s.Id == "" {
				s = riot.GetSummonerPuuid(summoner.Puuid)
				_, err := db.Db.Model(&s).Insert()
				if err != nil {
					fmt.Println("Error (saveMatch): ", err)
				}
			}
		} else {
			s = (*summoners)[0]
		}
		summonersMatch = append(summonersMatch, s)
	}

	return summonersMatch
}

func findSummoner(summoners *[]model.Summoner, summoner *model.Summoner) int {
	for i, s := range *summoners {
		if s.Puuid == summoner.Puuid {
			return i
		}
	}
	return -1
}

func saveMatch(match *model.Match) {
	_, err := db.Db.Model(match).Insert()
	if err != nil {
		fmt.Println("Error (saveMatch): ", err)
		return
	}
	for _, summoner := range match.Summoners {
		_, err = db.Db.Model(&model.MatchSummoner{MatchID: match.IdDB, SummonerID: summoner.Id}).Insert()
		if err != nil {
			fmt.Println("Error (saveMatch): ", err)
		}
	}
}
