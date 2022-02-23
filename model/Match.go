package model

import (
	"fmt"
	"strings"
	"time"
)

type Match struct {
	Id        string     `json:"matchId" pg:"match_id"`
	IdDB      int64      `pg:"id"`
	Start     time.Time  `pg:"start"`
	Summoners []Summoner `pg:"many2many:match_summoners"`
}

func (ms Match) String() string {
	return fmt.Sprintf(
		"stat: https://www.leagueofgraphs.com/match/%s/%s date: %s",
		strings.ToLower(strings.Split(ms.Id, "_")[0]),
		strings.Split(ms.Id, "_")[1],
		ms.Start,
	)
}
