package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"joint-games/command"
	"joint-games/riot"
	"log"
	"os"
)

func main() {

	var start, countGame int

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "global-region",
				Aliases:     []string{"gr"},
				Usage:       "global region (europe, americas, asia)",
				EnvVars:     []string{"GLOBAL_REGION"},
				DefaultText: os.Getenv("GLOBAL_REGION"),
				Destination: &riot.GlobalRegion,
			},
			&cli.StringFlag{
				Name:        "region",
				Aliases:     []string{"r"},
				Usage:       "server region (ru, br1, euw1, la1, ...)",
				EnvVars:     []string{"REGION"},
				DefaultText: os.Getenv("REGION"),
				Destination: &riot.Region,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "parser",
				Aliases: []string{"p"},
				Usage:   "Parses all the summoner's games",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "start",
						Aliases:     []string{"r"},
						Usage:       "Starting countdown for games",
						Value:       0,
						DefaultText: "0",
						Destination: &start,
					},
					&cli.BoolFlag{
						Name:    "update",
						Aliases: []string{"u"},
						Usage:   "Parse only the latest non-saved matches",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						fmt.Println("Specify the player's nickname as the first parameter")
						return nil
					}
					command.Parser(c.Args().First(), start, c.Bool("update"))
					return nil
				},
			},
			{
				Name:    "frequent",
				Aliases: []string{"f"},
				Usage:   "Getting Summoners you play with more often",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "count",
						Aliases:     []string{"c"},
						Usage:       "Minimum number of games with the summoner",
						Value:       4,
						DefaultText: "4",
						Destination: &countGame,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						fmt.Println("Specify the player's nickname as the first parameter")
						return nil
					}
					command.Frequent(c.Args().First(), countGame)
					return nil
				},
			},
			{
				Name:    "active",
				Aliases: []string{"a"},
				Usage:   "Checks an active match and finds past players",
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						fmt.Println("Specify the player's nickname as the first parameter")
						return nil
					}

					command.ActiveGame(c.Args().First())
					return nil
				},
			},
			{
				Name:    "jointMatches",
				Aliases: []string{"jm"},
				Usage:   "Returns joint games with the summoner",
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 2 {
						fmt.Println("It is necessary to specify two nicknames of the summoners")
						return nil
					}

					command.JointMatches(c.Args().Get(0), c.Args().Get(1))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
