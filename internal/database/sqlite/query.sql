-- name: GetAlbum :one
SELECT * FROM Album
WHERE AlbumId=?;

-- name: ListAlbums :many
SELECT * FROM Album;