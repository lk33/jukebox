SELECT ma."name"
FROM music_albums ma
INNER JOIN album_musicians am ON ma.album_id = am.album_id
INNER JOIN musicians m ON am.musician_id = m.musician_id
WHERE m.name = $1
ORDER BY ma.price ASC;
