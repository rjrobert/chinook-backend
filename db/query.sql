-- name: GetArtists :many
SELECT * FROM Artist;

-- name: GetArtist :one
SELECT * FROM Artist
WHERE ArtistId=?;

-- name: GetAllAlbums :many
SELECT * FROM Album;

-- name: GetAlbum :one
SELECT * FROM Album
WHERE AlbumId=?;

-- name: GetAlbumsByArtist :many
SELECT * FROM Album
WHERE ArtistId=?;

-- name: GetTracksByAlbum :many
SELECT t.TrackId, t.Name, mt.Name 'Media Type', g.Name 'Genre',
        t.Composer, t.Milliseconds, t.Bytes, t.UnitPrice
FROM Track t
JOIN MediaType mt ON mt.MediaTypeId=t.MediaTypeId
JOIN Genre g ON g.GenreId=t.GenreId
WHERE t.AlbumId=?;

-- name: GetPlaylists :many
SELECT * FROM Playlist;

-- name: GetPlaylistTracks :many
SELECT t.TrackId, t.Name, mt.Name 'Media Type', g.Name 'Genre',
        t.Composer, t.Milliseconds, t.Bytes, t.UnitPrice
FROM PlaylistTrack pt
JOIN Track t ON t.TrackId=pt.TrackId
JOIN MediaType mt ON mt.MediaTypeId=t.MediaTypeId
JOIN Genre g ON g.GenreId=t.GenreId
WHERE pt.PlaylistId=?;