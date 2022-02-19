package model

type MatchSummoner struct {
	MatchID    int       `pg:",pk"`
	Match      *Match    `pg:"rel:has-one"`
	SummonerID int       `pg:",pk"`
	Summoner   *Summoner `pg:"rel:has-one"`
}
