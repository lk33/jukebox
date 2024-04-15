-- Create musicians table
CREATE TABLE musicians (
    musician_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    musician_type VARCHAR(50),
    CONSTRAINT unique_name_musician_type UNIQUE ("name", musician_type);
);

-- Create music_albums table
CREATE TABLE music_albums (
    album_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    genre VARCHAR(100),
    price DECIMAL(10, 2) NOT NULL,
    description TEXT,
    CONSTRAINT unique_album_name UNIQUE (name)
);

-- Create album_musicians junction table
CREATE TABLE album_musicians (
    album_id INT NOT NULL,
    musician_id INT NOT NULL,
    musician_type VARCHAR(50),
    PRIMARY KEY (album_id, musician_id),
    FOREIGN KEY (album_id) REFERENCES music_albums(album_id) ON DELETE CASCADE,
    FOREIGN KEY (musician_id) REFERENCES musicians(musician_id) ON DELETE CASCADE
);
