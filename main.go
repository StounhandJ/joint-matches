package main

import (
	"joint-games/riot"
)

func main() {
	//db := database.NewDataBase()
	riot.GetMatches(riot.GetSummoner("StounhandJ"))

}
