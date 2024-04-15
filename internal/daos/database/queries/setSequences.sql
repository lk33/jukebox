
CREATE SEQUENCE album_id_seq;

ALTER TABLE music_albums
    ALTER COLUMN album_id SET DEFAULT nextval('album_id_seq');
	
CREATE SEQUENCE music_id_seq;

ALTER TABLE musicians
    ALTER COLUMN musician_id SET DEFAULT nextval('music_id_seq');