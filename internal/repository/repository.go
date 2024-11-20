package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"

	"github.com/rjrobert/chinook-music/backend/internal/repository/database"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	GetArtists(ctx context.Context) ([]database.Artist, error)
	GetArtist(ctx context.Context, artistId int64) (*database.Artist, error)
	GetAllAlbums(ctx context.Context) ([]database.Album, error)
	GetAlbum(ctx context.Context, albumId int64) (*database.Album, error)
	GetAlbumsByArtist(ctx context.Context, artistId int64) ([]database.Album, error)
	GetTracksByAlbum(ctx context.Context, albumId int64) ([]database.GetTracksByAlbumRow, error)
	GetPlaylists(ctx context.Context) ([]database.Playlist, error)
	GetPlayListTracks(ctx context.Context, playlistId int64) ([]database.GetPlaylistTracksRow, error)
}

type service struct {
	db      *sql.DB
	queries *database.Queries
}

var (
	dburl      = os.Getenv("SQLITE_DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	queries := database.New(db)

	dbInstance = &service{
		db:      db,
		queries: queries,
	}
	return dbInstance
}

func (s *service) GetArtists(ctx context.Context) ([]database.Artist, error) {
	artists, err := s.queries.GetArtists(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting albums: %w", err)
	}
	return artists, nil
}

func (s *service) GetArtist(ctx context.Context, artistId int64) (*database.Artist, error) {
	artist, err := s.queries.GetArtist(ctx, artistId)
	if err != nil {
		return nil, fmt.Errorf("getting artist: %w", err)
	}
	return &artist, nil
}

func (s *service) GetAllAlbums(ctx context.Context) ([]database.Album, error) {
	albums, err := s.queries.GetAllAlbums(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting albums: %w", err)
	}
	return albums, nil
}

func (s *service) GetAlbum(ctx context.Context, albumId int64) (*database.Album, error) {
	album, err := s.queries.GetAlbum(ctx, albumId)
	if err != nil {
		return nil, fmt.Errorf("getting album: %w", err)
	}
	return &album, nil
}

func (s *service) GetAlbumsByArtist(ctx context.Context, artistId int64) ([]database.Album, error) {
	albums, err := s.queries.GetAlbumsByArtist(ctx, artistId)
	if err != nil {
		return nil, fmt.Errorf("getting albums by artist: %w", err)
	}
	return albums, nil
}

func (s *service) GetTracksByAlbum(ctx context.Context, albumId int64) ([]database.GetTracksByAlbumRow, error) {
	tracks, err := s.queries.GetTracksByAlbum(ctx, &albumId)
	if err != nil {
		return nil, fmt.Errorf("getting tracks by album: %w", err)
	}
	return tracks, nil
}

func (s *service) GetPlaylists(ctx context.Context) ([]database.Playlist, error) {
	playlists, err := s.queries.GetPlaylists(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting playlists: %w", err)
	}
	return playlists, nil
}

func (s *service) GetPlayListTracks(ctx context.Context, playlistId int64) ([]database.GetPlaylistTracksRow, error) {
	tracks, err := s.queries.GetPlaylistTracks(ctx, playlistId)
	if err != nil {
		return nil, fmt.Errorf("getting playlist tracks: %w", err)
	}
	return tracks, nil
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}
