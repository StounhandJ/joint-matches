package riot

import (
	"encoding/json"
	"errors"
	"fmt"
	"joint-games/model"
	"strconv"
	"time"
)

func GetSummoner(name string) model.Summoner {
	var summoner model.Summoner

	_ = json.
		NewDecoder(Get(Region, fmt.Sprintf("lol/summoner/v4/summoners/by-name/%s", name), nil)).
		Decode(&summoner)

	return summoner
}

func GetMatches(summoner model.Summoner, start int, startTime int64) chan model.Match {
	ch := make(chan model.Match)

	go func() {
		defer close(ch)
		work := true
		i := start / 100

		for work {

			matchesId := getMatchesIds(summoner, i*100, startTime)

			for _, id := range matchesId {
				match, err := GetMatch(id)
				if err != nil {
					fmt.Println(fmt.Sprintf("Error (GetMatches) %s: %s", id, err))
					continue
				}
				ch <- match
			}

			work = len(matchesId) == 100
			i += 1
		}
	}()

	return ch
}

func getMatchesIds(summoner model.Summoner, start int, startTime int64) []string {
	var ids []string

	_ = json.
		NewDecoder(Get(
			GlobalRegion,
			fmt.Sprintf("lol/match/v5/matches/by-puuid/%s/ids", summoner.Puuid),
			map[string]string{"count": "100", "start": strconv.Itoa(start), "startTime": strconv.FormatInt(startTime, 10)})).
		Decode(&ids)

	return ids
}

func GetMatch(id string) (model.Match, error) {

	var matchData map[string]interface{}
	match := model.Match{Id: id}

	_ = json.
		NewDecoder(Get(
			GlobalRegion,
			fmt.Sprintf("lol/match/v5/matches/%s", id),
			nil)).
		Decode(&matchData)

	if _, ok := matchData["status"]; ok {
		return match, errors.New("match not found")
	}

	for _, puuid := range matchData["metadata"].(map[string]interface{})["participants"].([]interface{}) {
		match.Summoners = append(match.Summoners, model.Summoner{Puuid: puuid.(string)})
	}
	match.Start = time.Unix(int64(matchData["info"].(map[string]interface{})["gameCreation"].(float64))/1000, 0)

	return match, nil
}

func GetActiveMatch(summonerId string) (model.Match, error) {

	var matchData map[string]interface{}
	match := model.Match{}

	_ = json.
		NewDecoder(Get(
			Region,
			fmt.Sprintf("lol/spectator/v4/active-games/by-summoner/%s", summonerId),
			nil)).
		Decode(&matchData)

	if _, ok := matchData["status"]; ok {
		return match, errors.New("match not found")
	}

	for _, participantsData := range matchData["participants"].([]interface{}) {
		if !participantsData.(map[string]interface{})["bot"].(bool) {
			match.Summoners = append(match.Summoners, model.Summoner{
				Id:   participantsData.(map[string]interface{})["summonerId"].(string),
				Name: participantsData.(map[string]interface{})["summonerName"].(string)})
		}
	}

	return match, nil
}

func GetSummonerPuuid(puuid string) model.Summoner {
	var summoner model.Summoner

	_ = json.
		NewDecoder(Get(Region, fmt.Sprintf("lol/summoner/v4/summoners/by-puuid/%s", puuid), nil)).
		Decode(&summoner)
	return summoner
}
