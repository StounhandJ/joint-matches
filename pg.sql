create table "match"(
	"id" serial PRIMARY KEY,
	"match_id" varchar(15)
);

create table "summoner"(
	"id" varchar(80) PRIMARY KEY,
	"accountId" varchar(80),
	"puuid" varchar(120)
);

create table "match_summoner"(
	"id_match" int REFERENCES "match" ("id"),
	"id_summoner" varchar(80) REFERENCES "summoner" ("id")
);
