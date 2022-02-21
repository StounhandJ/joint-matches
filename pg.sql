create table "matches"(
	"id" serial PRIMARY KEY,
	"match_id" varchar(15),
	"start" timestamp
);

create table "summoners"(
	"id" varchar(80) PRIMARY KEY,
	"accountId" varchar(80),
	"puuid" varchar(120),
	"name" varchar(120)
);

create table "match_summoners"(
	"match_id" int REFERENCES "matches" ("id"),
	"summoner_id" varchar(80) REFERENCES "summoners" ("id")
);
