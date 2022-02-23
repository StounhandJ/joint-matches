package database

import "joint-matches/model"

func LastMatch(summoner model.Summoner) (*model.Match, error) {

	db := NewDataBase()

	var lastMatch model.Match

	_, err := db.Db.Query(
		&lastMatch,
		"select matches.id, matches.start, matches.match_id from matches "+
			"join match_summoners on match_summoners.match_id = matches.id "+
			"where match_summoners.summoner_id = ? "+
			"order by \"start\" DESC "+
			"Limit 1", summoner.Id)

	if err != nil {
		return nil, err
	}

	return &lastMatch, nil
}
