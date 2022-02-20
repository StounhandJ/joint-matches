package riot

import (
	"encoding/json"
	"fmt"
	"joint-games/model"
	"strconv"
)

func GetSummoner(name string) model.Summoner {
	var summoner model.Summoner

	_ = json.
		NewDecoder(Get("ru", fmt.Sprintf("lol/summoner/v4/summoners/by-name/%s", name), nil)).
		Decode(&summoner)
	return summoner
}

func GetMatches(summoner model.Summoner) []model.Match {

	var matches []model.Match

	work := true
	i := 1

	for work {

		matchesId := getMatchesIds(summoner, i*100)

		for _, id := range matchesId {
			matches = append(matches, GetMatch(id))
		}

		work = len(matchesId) == 100
		i += 1
	}

	return matches
}

func getMatchesIds(summoner model.Summoner, start int) []string {
	var ids []string

	_ = json.
		NewDecoder(Get(
			"europe",
			fmt.Sprintf("lol/match/v5/matches/by-puuid/%s/ids", summoner.Puuid),
			map[string]string{"count": "100", "start": strconv.Itoa(start)})).
		Decode(&ids)

	return ids
}

func GetMatch(id string) model.Match {

	var matchData map[string]interface{}
	match := model.Match{Id: id}

	_ = json.
		NewDecoder(Get(
			"europe",
			fmt.Sprintf("lol/match/v5/matches/%s", id),
			nil)).
		Decode(&matchData)

	for _, puuid := range matchData["metadata"].(map[string]interface{})["participants"].([]interface{}) {
		match.Summoners = append(match.Summoners, model.Summoner{Puuid: puuid.(string)})
	}

	return match
}
