
--- Create Users for the App
CREATE USER Was_Owner WITH PASSWORD = ''
CREATE USER Was_Writer WITH PASSWORD = '';
CREATE USER Was_Reader WITH PASSWORD = '';
GO
--- Create SCHEMA for the database as well as tables
CREATE SCHEMA was_groupwork AUTHORIZATION Was_Owner
	CREATE TABLE [User] (
		id int Identity(1000,1) PRIMARY KEY,
		name varchar(60) NOT NULL,
		email varchar(260) UNIQUE NOT NULL
	)
	CREATE TABLE BusinessAnnouncement (
		id int Identity(100,1) PRIMARY KEY,
		[user_id] int NOT NULL,
		expiration_date datetime NOT NULL,
		category varchar(8) NOT NULL CHECK(category IN('buying', 'selling')),
		announcement nvarchar(max)
	)
GO
ALTER TABLE was_groupwork.BusinessAnnouncement 
	ADD CONSTRAINT FK_BusinessAnnouncement_User FOREIGN KEY ([user_id])
		REFERENCES was_groupwork.[User] (id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
;
GO
