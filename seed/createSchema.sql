-- creating schema
-- User table with name email and password
CREATE TABLE IF NOT EXISTS userinfo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Book table with Authorname,Subjectname,Semnumber,Branch,Universityname,Bookfile and Bannerimage

CREATE TABLE IF NOT EXISTS bookinfo (
    id SERIAL PRIMARY KEY,
    bookauthor VARCHAR(255) NOT NULL,
    subjectname VARCHAR(255) NOT NULL,
    semester VARCHAR(255) NOT NULL,
    branch VARCHAR(255) NOT NULL,
    universityname VARCHAR(255) NOT NULL,
    bookpath VARCHAR(255) NOT NULL,
    imgpath VARCHAR(255) NOT NULL
);
          
CREATE TABLE IF NOT EXISTS Role(
RoleID SERIAL PRIMARY KEY,
RoleName varchar(50)
);    

insert into Role (RoleName)
values ('Admin');