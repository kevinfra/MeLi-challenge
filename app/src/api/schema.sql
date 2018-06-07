-- Create items table
CREATE TABLE IF NOT EXISTS items (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name TEXT,
	description TEXT
);

-- Create documents table
CREATE TABLE IF NOT EXISTS documents (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	driveId TEXT,
	title TEXT,
	description TEXT
);

-- Create tokens table
CREATE TABLE IF NOT EXISTS tokens (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	token TEXT
);