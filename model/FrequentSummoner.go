package model

type FrequentSummoner struct {
	Id    string `pg:"id"`
	Puuid string `pg:"puuid"`
	Name  string `pg:"name"`
	Count int    `pg:"count"`
}
