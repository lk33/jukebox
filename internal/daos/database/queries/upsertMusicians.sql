INSERT INTO musicians (musician_id, "name", musician_type)
VALUES (DEFAULT, $1, $2)
ON CONFLICT ("name") DO UPDATE
SET 
    musician_type = EXCLUDED.musician_type;