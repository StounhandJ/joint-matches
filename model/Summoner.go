package model

type Summoner struct {
	Id        string `json:"id" pg:"id"`
	AccountId string `json:"accountId" pg:"accountId"`
	Puuid     string `json:"puuid" pg:"puuid"`
	Name      string `json:"name" pg:"name"`
}
