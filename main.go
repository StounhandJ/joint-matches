package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"joint-games/command"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Joint games for League of Legends")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "help":
		command.Help()
		break
	case "parser":
		command.Parser()
	case "frequent":
		command.Frequent()
	default:
		command.Help()
		break
	}
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
