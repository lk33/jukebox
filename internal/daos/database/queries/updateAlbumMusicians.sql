INSERT INTO album_musicians (album_id, musician_id, musician_type)
VALUES ($1, $2, $3)
ON CONFLICT (album_id, musician_id, musician_type) DO UPDATE
SET musician_type = EXCLUDED.musician_type
RETURNING *;