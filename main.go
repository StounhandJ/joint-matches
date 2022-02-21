package main

import (
	"fmt"
	"joint-games/database"
	"joint-games/model"
	"joint-games/riot"
)

var (
	db database.DataBase
)

func main() {

	db = *database.NewDataBase()
	db.Db.Model(&model.MatchSummoner{})
	var summoners []model.Summoner
	i := 0

	for match := range riot.GetMatches(riot.GetSummoner("StounhandJ")) {
		fmt.Println(i)
		i += 1

		_ = db.Db.Model(&match).Where("match_id = ?", match.Id).Select()
		if match.IdDB != 0 {
			continue
		}

		match.Summoners = getSummoners(&match.Summoners, &summoners)
		copy(summoners, match.Summoners[:])

		saveMatch(&match)
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
				_, _ = db.Db.Model(&s).Insert()
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
	_, _ = db.Db.Model(&match).Insert()
	for _, summoner := range match.Summoners {
		_, _ = db.Db.Model(&model.MatchSummoner{MatchID: match.IdDB, SummonerID: summoner.Id}).Insert()
	}
}
