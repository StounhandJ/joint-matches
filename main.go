package main

import (
	"awesomeProject2/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//message := map[string]interface{}{
	//	"hello": "world",
	//	"life":  42,
	//	"embedded": map[string]string{
	//		"yes": "of course!",
	//	},
	//}
	//
	//bytesRepresentation, err := json.Marshal(message)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//resp, err := http.Get("https://ru.api.riotgames.com/lol/summoner/v4/summoners/by-name/StounhandJ?api_key=RGAPI-e8de3f30-d68c-469f-b489-aa96f2674d42")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//var result map[string]interface{}
	//
	//json.NewDecoder(resp.Body).Decode(&result)
	//
	//log.Println(result)

	GetMatches(GetSummoner("StounhandJ"))
}

func GetSummoner(name string) model.Summoner {
	req, err := http.NewRequest("GET", "https://ru.api.riotgames.com/lol/summoner/v4/summoners/by-name/"+name, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "RGAPI-e8de3f30-d68c-469f-b489-aa96f2674d42")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var summoner model.Summoner

	_ = json.NewDecoder(resp.Body).Decode(&summoner)
	return summoner
}

func GetMatches(summoner model.Summoner) []model.Match {

	var matches []model.Match

	re := true
	i := 1

	for re {
		var arr []string

		_ = json.NewDecoder(getMatchesIds(summoner, i*100)).Decode(&arr)

		for _, element := range arr {
			matches = append(matches, model.Match{Id: element})
		}

		re = len(arr) == 100
		i += 1
	}

	return matches
}

func getMatchesIds(summoner model.Summoner, start int) io.Reader {
	req, err := http.NewRequest("GET", "https://europe.api.riotgames.com/lol/match/v5/matches/by-puuid/"+summoner.Puuid+"/ids", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "RGAPI-e8de3f30-d68c-469f-b489-aa96f2674d42")
	q.Add("count", "100")
	q.Add("start", strconv.Itoa(start))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp.Body
}
