package database

import (
	"joint-matches/model"
)

func FrequentSummonerInMatch(summoner model.Summoner, countGame int) ([]model.FrequentSummoner, error) {

	db := NewDataBase()

	var frequentSummoners []model.FrequentSummoner

	_, err := db.Db.Query(
		&frequentSummoners,
		"select summ.id, summ.puuid, summ.name, COUNT(*) from match_summoners as main "+
			"join match_summoners as dop on main.match_id = dop.match_id and main.summoner_id != ? "+
			"join summoners as summ on summ.id = main.summoner_id "+
			"where dop.summoner_id = ? "+
			"GROUP BY summ.id HAVING COUNT(*) >= ? "+
			"ORDER BY \"count\" DESC;", summoner.Id, summoner.Id, countGame)

	if err != nil {
		return nil, err
	}

	return frequentSummoners, nil
}
