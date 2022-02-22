package command

import "fmt"

func Help() {
	fmt.Println(
		"parser - Parses all the summoner's games\n\t1 - Summoner's nickname\n\t2 - Starting countdown for games\n" +
			"frequent - Getting Summoners you play with more often\n\t1 - Summoner's nickname\n\t2 - Minimum number of games with the summoner")
}
