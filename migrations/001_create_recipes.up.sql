CREATE TABLE IF NOT EXISTS recipes (
		id                   bigserial PRIMARY KEY,
    name                 varchar(256) NOT NULL,
    prep_time_in_minutes integer NOT NULL,
    difficulty           integer NOT NULL,
    vegetarian           boolean NOT NULL
);
