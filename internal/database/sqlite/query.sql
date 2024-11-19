-- name: GetAlbum :one
SELECT * FROM Album
WHERE AlbumId=?;

-- name: GetAlbums :many
SELECT * FROM Album;