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

-- name: GetGenres :many
SELECT *
FROM Genre;

-- name: GetMediaTypes :many
SELECT * 
FROM MediaType;

-- name: GetCustomers :many
SELECT *
FROM Customer;

-- name: CreateCustomer :one
INSERT INTO Customer (FirstName, LastName, Company, Address, City, State, Country, PostalCode, Phone, Fax, Email, SupportRepId)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetCustomerInvoices :many
SELECT *
FROM Invoice
WHERE CustomerId=?;

-- name: GetInvoiceLines :many
SELECT *
FROM InvoiceLine
WHERE InvoiceId=?;
