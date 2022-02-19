create table "matches"(
	"id" serial PRIMARY KEY,
	"match_id" varchar(15)
);

create table "summoners"(
	"id" varchar(80) PRIMARY KEY,
	"accountId" varchar(80),
	"puuid" varchar(120)
);

create table "match_summoners"(
	"id_match" int REFERENCES "matches" ("id"),
	"id_summoner" varchar(80) REFERENCES "summoners" ("id")
);
