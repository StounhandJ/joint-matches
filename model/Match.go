package model

type Match struct {
	Id   string `json:"matchId" pg:"matchId"`
	IdDB string `pg:"id"`
}
