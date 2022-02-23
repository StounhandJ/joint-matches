package model

import (
	"fmt"
)

type Summoner struct {
	Id        string `json:"id" pg:"id"`
	AccountId string `json:"accountId" pg:"accountId"`
	Puuid     string `json:"puuid" pg:"puuid"`
	Name      string `json:"name" pg:"name"`
}

func (s Summoner) String() string {
	return fmt.Sprintf(
		"NickName: %s",
		s.Name,
	)
}
