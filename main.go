package main

import (
	"awesomeProject2/database"
	"awesomeProject2/model"
)

func main() {
	//riot.GetMatches(riot.GetSummoner("StounhandJ"))
	db := database.NewDataBase()
	//s1 := &model.Summoner{AccountId: "ghjghj", Puuid: "fghfghfgh", Id: "hhjhgjghj"}
	//s2 := &model.Summoner{AccountId: "hkjlh", Puuid: "lklkjjklj", Id: "lllllll"}
	//db.Insert(s1)
	//db.Insert(s2)
	db.Db.Model(&model.MatchSummoner{})
	var summoners []model.Summoner
	err := db.Db.Model(&summoners).Select()
	if err != nil {
		panic(err)
	}
	db.Db.Model(&model.Match{Id: "ppppppppppp", Summoners: summoners}).Insert()
}
