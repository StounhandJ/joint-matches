-- Get the most games with another player
select summ.id, summ.name, COUNT(*) from match_summoners as main
	join match_summoners as dop on main.match_id = dop.match_id and main.summoner_id != ''
	join summoners as summ on summ.id = main.summoner_id
	where dop.summoner_id = ''
	GROUP BY summ.id
	HAVING COUNT(*) >= 3
	ORDER BY "count" DESC;

-- Get a list of games with a specific player
select matches.match_id, matches.start from match_summoners as main
	join match_summoners as dop on main.match_id = dop.match_id
	join matches on matches.id = main.match_id
	where dop.summoner_id = '' and main.summoner_id =''