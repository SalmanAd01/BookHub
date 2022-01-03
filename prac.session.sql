SELECT *
from movies;
CREATE TABLE userinfo (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL
)
SELECT *
FROM PG_CATALOG.pg_tables
WHERE schemaname = 'public' DROP TABLE usser
SELECT *
From userinfo;
DROP TABLE userinfo -- sdfdsf