package model

import "time"

type Match struct {
	Id        string     `json:"matchId" pg:"match_id"`
	IdDB      int64      `pg:"id"`
	Start     time.Time  `pg:"start"`
	Summoners []Summoner `pg:"many2many:match_summoners"`
}
