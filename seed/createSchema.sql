-- creating schema

-- User table
CREATE TABLE IF NOT EXISTS userinfo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Book table
CREATE TABLE IF NOT EXISTS bookinfo (
    id SERIAL PRIMARY KEY,
    bookpath VARCHAR(255) NOT NULL,
    imgpath VARCHAR(255) NOT NULL,
    subjectname VARCHAR(255) NOT NULL,
    bookauthor VARCHAR(255) NOT NULL userinfo(id),
    semester VARCHAR(255) NOT NULL,
    branch VARCHAR(255) NOT NULL,
    universityname VARCHAR(255) NOT NULL,
    userid FOREIGN KEY REFERENCES userinfo(id)
);