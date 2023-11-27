-- Create Users schema

CREATE TABLE IF NOT EXISTS Users (
	uid INTEGER  PRIMARY KEY, 	-- user id
	username TEXT UNIQUE NOT NULL	-- username of user
);

-- Create Photo schema

CREATE TABLE IF NOT EXISTS Photos (
	photoId INTEGER  PRIMARY KEY,	-- photo identificator
	owner INTEGER NOT NULL,		-- uid of owner
	descriptionImg TEXT NOT NULL,	-- image description
	ImageData BLOB NOT NULL,	-- file
	timeUpdate TIMESTAMP NOT NULL,	-- time to update
	
	
	CONSTRAINT ref_to_user FOREIGN KEY(owner) REFERENCES Users(uid)
);


-- Create Comments schema

CREATE TABLE IF NOT EXISTS Comments (
	commentId INTEGER   PRIMARY KEY,
	author INTEGER NOT NULL,
	photo INTEGER NOT NULL,
	text_ TEXT NOT NULL,
	timeStamp_ TIMESTAMP NOT NULL,

	CONSTRAINT ref_to_user FOREIGN KEY(author) REFERENCES Users(uid),
	CONSTRAINT ref_to_photo FOREIGN KEY(photo) REFERENCES Photos(photoId)
);


-- Create likes schema

CREATE TABLE IF NOT EXISTS Likes (
	user INTEGER NOT NULL,
	photo INTEGER NOT NULL,
	PRIMARY KEY(user, photo),
	CONSTRAINT ref_to_user FOREIGN KEY(user) REFERENCES Users(uid),
	CONSTRAINT ref_to_photo FOREIGN KEY(photo) REFERENCES Photos(photoId)
);


-- Create followers schema

CREATE TABLE IF NOT EXISTS Followers (
	from_ INTEGER NOT NULL,
	to_ INTEGER NOT NULL,
	
	PRIMARY KEY(from_, to_)
	CONSTRAINT ref_to_from FOREIGN KEY(from_) REFERENCES User(uid),
	CONSTRAINT ref_to FOREIGN KEY(to_) REFERENCES User(uid)
);
-- Create banned schema

CREATE TABLE IF NOT EXISTS Bans (
	from_ INTEGER NOT NULL,
	to_ INTEGER NOT NULL,
	
	PRIMARY KEY(from_, to_),
	CONSTRAINT ref_to_from FOREIGN KEY(from_) REFERENCES User(uid)
	CONSTRAINT ref_to FOREIGN KEY(to_) REFERENCES User(uid)
);
