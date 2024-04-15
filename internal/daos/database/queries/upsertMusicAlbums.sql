INSERT INTO music_albums (album_id, "name", release_date, genre, price, description)
VALUES (DEFAULT, $1, $2, $3, $4, $5)
ON CONFLICT ("name") DO UPDATE
SET 
    release_date = EXCLUDED.release_date,
    genre = EXCLUDED.genre,
    price = EXCLUDED.price,
    "description" = EXCLUDED.description;