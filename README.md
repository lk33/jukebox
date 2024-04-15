# jukebox
Discover and manage music albums and artists with this JukeBox application. Built with Go, this project offers intuitive RESTful APIs for seamless retrieval and addition of music albums and artist details, enriching your music experience.

Trigger VSCode debugger to execute
APIs:
1. UpsertAlbum 
POST http://localhost:8080/jukebox/v1/album
2. UpsertMusician
POST http://localhost:8080/jukebox/v1/musician
3. GetAlbums (sorted by album release date) 
POST http://localhost:8080/jukebox/v1/albums
4. GetAlbumsByMusicianName (sorted by album price)
GET http://localhost:8080/jukebox/v1/albums?musician_name=Chester%20Bennington
5. GetMusiciansByAlbum (sorted by musician name)
GET http://localhost:8080/jukebox/v1/musicians?album_name=Meteora
