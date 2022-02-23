Joint Matches - Анализ матчей LoL, чтобы найти общие игры с другими игроками
===
[[EN]](README.md) [RU]
### При активной игре будет отображаться такой список призывателей из вашей игры.
```shell
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/377256662 date: 2022-02-23 16:14:47 +0000 UTC
NickName: кiшка | stat: https://www.leagueofgraphs.com/match/ru/357069038 date: 2021-10-31 21:25:30 +0000 UTC
NickName: LeAnk | stat: https://www.leagueofgraphs.com/match/ru/343836977 date: 2021-07-30 14:03:37 +0000 UTC
```
### По стандарту работает с сервером "ru", но используя параметры окружение и options можно указать любой регион.

---

КОМАНДЫ:
---
- [**parser**](#-parser), p - Анализирует все игры призывателя
- [**frequent**](#-frequent), f - Получение призывателей, с которыми вы играете чаще
- [**active**](#-active), a - Проверяет активный матч и находит прошлых игроков
- [**jointMatches**](#-jointMatches), jm - Возвращает совместные игры с призывателем

ГЛОБАЛЬНЫЕ ПАРАМЕТРЫ:
---
- **--global-region value**, **--gr** *value* - глобальный регион (europe, americas, asia) (de
fault: europe)
- **--region value**, **-r** *value* - регион сервера (ru, br1, euw1, la1, ...) (d
efault: ru)

---

## • active
Проверяет активный матч и находит прошлых игроков
#### ARGUMENTS:
- Никнейм призывателя
#### EXAMPLES:
1. Получите список призывателей из активной игры со ссылкой на старые матчи
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
Возвращает совместные игры с призывателем
#### ARGUMENTS:
- Никнейм призывателя
- Никнейм второго призывателя
#### EXAMPLES:
1. Получить список игр с данным призывателем
```shell
./joint-matches jointMatches StounhandJ кiшка
```
Результат: 
```shell
stat: https://www.leagueofgraphs.com/match/ru/358378025 date: 2021-11-08 19:23:23 +0000 UTC
stat: https://www.leagueofgraphs.com/match/ru/357069038 date: 2021-10-31 21:25:30 +0000 UTC
stat: https://www.leagueofgraphs.com/match/ru/377256662 date: 2022-02-23 16:14:47 +0000 UTC
```

---

## • parser
Анализирует все игры призывателя
#### ARGUMENTS:
- Никнейм призывателя
#### OPTIONS:
   --start value, -r value - Количество пропускаемых матчей от новых к старым (default: 0)  
   --update, -u            - Анализировать только последние несохраненные совпадения (default: false)
#### EXAMPLES:
1. Получите все игры призывателя "StounhandJ"
```shell
./joint-matches parser --start=0 StounhandJ
```
2. Получите все игры призывателя "StounhandJ", пропустив последние 200
```shell
./joint-matches parser --start=200 StounhandJ
```
3. Получите последние несохраненные матчи призывателя "StounhandJ"
```shell
./joint-matches parser --update StounhandJ
```

---

## • frequent
Получение призывателей, с которыми вы играете чаще
#### ARGUMENTS:
- Никнейм призывателя
#### OPTIONS:
   --count value, -c value - Минимальное количество игр с призывателем (default: 4)
#### EXAMPLES:
1. Получить призывателей, с которыми "StounhandJ" вместе сыграли более 4 игр 
```shell
./joint-matches frequent StounhandJ
```
2. Получить призывателей, с которыми "StounhandJ" вместе сыграли более 10 игр
```shell
./joint-matches frequent --count=10 StounhandJ
```