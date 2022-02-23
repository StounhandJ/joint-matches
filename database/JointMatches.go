package database

import (
	"joint-games/model"
)

func JointMatches(summonerOne model.Summoner, summonerTwo model.Summoner) ([]model.Match, error) {
	db := NewDataBase()
	db.Db.Model(&model.MatchSummoner{})

	var matches []model.Match
	_, err := db.Db.Query(
		&matches,
		"select matches.id, matches.match_id, matches.start "+
			"from match_summoners as main "+
			"join match_summoners as dop on main.match_id = dop.match_id "+
			"join matches on matches.id = main.match_id "+
			"where dop.summoner_id = ? and main.summoner_id = ?", summonerOne.Id, summonerTwo.Id)
	if err != nil {
		return nil, err
	}
	return matches, nil
}
