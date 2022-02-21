package model

type MatchSummoner struct {
	MatchID    int64     `pg:",pk"`
	Match      *Match    `pg:"rel:has-one"`
	SummonerID string    `pg:",pk"`
	Summoner   *Summoner `pg:"rel:has-one"`
}
