-- Get the most games with another player
select main.summoner_id, COUNT(*) from match_summoners as main
	join match_summoners as dop on main.match_id = dop.match_id where dop.summoner_id = ''
	GROUP BY main.summoner_id
	ORDER BY "count" DESC;

-- Get a list of games with a specific player
select matches.match_id, matches.start from match_summoners as main
	join match_summoners as dop on main.match_id = dop.match_id
	join matches on matches.id = main.match_id
	where dop.summoner_id = '' and main.summoner_id =''