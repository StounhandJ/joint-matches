Joint Matches - LoL match analysis to find common games with other players
===
[EN] [[RU]](README_RU.md)
### With an active game, such a list of summoners from your game will be displayed.
```shell
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/377256662 date: 2022-02-23 16:14:47 +0000 UTC
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/357069038 date: 2021-10-31 21:25:30 +0000 UTC
NickName: LeAnk | stat: https://www.leagueofgraphs.com/match/ru/343836977 date: 2021-07-30 14:03:37 +0000 UTC
```
### According to the standard, it works with the "ru" server, but using the environment and options parameters, you can specify any region.  

---

COMMANDS:
---
- [**parser**](#-parser), p - Parses all the summoner's matches
- [**frequent**](#-frequent), f - Getting Summoners you play with more often
- [**active**](#-active), a - Checks an active match and finds past players
- [**jointMatches**](#-jointMatches), jm - Returns joint games with the summoner  

GLOBAL OPTIONS:
---
- **--global-region value**, **--gr** *value* - global region (europe, americas, asia) (de
fault: europe)
- **--region value**, **-r** *value* - server region (ru, br1, euw1, la1, ...) (d
efault: ru)

---
  
## • active
Checks an active match and finds past players
#### ARGUMENTS:
- Summoner's nickname
#### EXAMPLES:
1. Get a list of summoners from an active game with a link to old matches
```shell
./joint-matches active StounhandJ
```
Result: 
```shell
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/377256662 date: 2022-02-23 16:14:47 +0000 UTC
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/357069038 date: 2021-10-31 21:25:30 +0000 UTC
NickName: LeAnk | stat: https://www.leagueofgraphs.com/match/ru/343836977 date: 2021-07-30 14:03:37 +0000 UTC
```

---

## • jointMatches
Returns joint games with the summoner
#### ARGUMENTS:
- Summoner's nickname
- Summoner's second nickname
#### EXAMPLES:
1. Get a list of games with this summoners
```shell
./joint-matches jointMatches StounhandJ кiшка
```
Result: 
```shell
stat: https://www.leagueofgraphs.com/match/ru/358378025 date: 2021-11-08 19:23:23 +0000 UTC
stat: https://www.leagueofgraphs.com/match/ru/357069038 date: 2021-10-31 21:25:30 +0000 UTC
stat: https://www.leagueofgraphs.com/match/ru/377256662 date: 2022-02-23 16:14:47 +0000 UTC
```

---

## • parser
#### ARGUMENTS:
- Summoner's nickname
#### OPTIONS:
   --start value, -r value - The number of missed matches from new to old (default: 0)  
   --update, -u            - Parse only the latest non-saved matches (default: false)
#### EXAMPLES:
1. Get all the summoner games "StounhandJ"
```shell
./joint-matches parser --start=0 StounhandJ
```
2. Get all the summoner games by skipping the last 200 "StounhandJ"
```shell
./joint-matches parser --start=200 StounhandJ
```
3. Get the latest unsaved matches of the summoner "StounhandJ"
```shell
./joint-matches parser --update StounhandJ
```

---

## • frequent
Getting Summoners you play with more often
#### ARGUMENTS:
- Summoner's nickname
#### OPTIONS:
   --count value, -c value - Minimum number of games with the summoner (default: 4)
#### EXAMPLES:
1. Get summoners with whom "StounhandJ" have played more than 4 games together
```shell
./joint-matches frequent StounhandJ
```
2. Get summoners with whom "StounhandJ" have played more than 10 games together
```shell
./joint-matches frequent --count=10 StounhandJ
```