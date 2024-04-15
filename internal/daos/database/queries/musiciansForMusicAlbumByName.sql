SELECT m."name", m.musician_type
FROM musicians m
INNER JOIN album_musicians am on m.musician_id = am.musician_id
INNER JOIN music_albums ma on am.album_id = ma.album_id
WHERE ma."name" = $1
ORDER BY m."name" ASC