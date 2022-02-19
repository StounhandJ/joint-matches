package model

type Match struct {
	Id        string     `json:"matchId" pg:"match_id"`
	IdDB      int64      `pg:"id"`
	Summoners []Summoner `pg:"many2many:match_summoners"`
}
