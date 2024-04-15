INSERT INTO musicians (musician_id, "name", musician_type)
VALUES (DEFAULT, $1, $2)
ON CONFLICT ("name", musician_type) DO NOTHING
RETURNING musician_id;
