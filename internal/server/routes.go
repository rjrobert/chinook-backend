package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", s.healthHandler)

	e.GET("/artists", s.getArtistsHandler)
	e.GET("/artists/:artistId", s.getArtistHandler)
	e.GET("/artists/:artistId/albums", s.getAlbumsByArtistHandler)

	e.GET("/albums", s.getAllAlbumsHandler)
	e.GET("/albums/:albumId", s.getAlbumHandler)
	e.GET("/albums/:albumId/tracks", s.getTracksByAlbumHandler)

	e.GET("/playlists", s.getPlaylistsHandler)
	e.GET("/playlists/:playlistId", s.getPlaylistTracksHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getArtistsHandler(c echo.Context) error {
	artists, err := s.db.GetArtists(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, artists)
}

func (s *Server) getArtistHandler(c echo.Context) error {
	artistIdStr := c.Param("artistId")
	artistId, err := strconv.ParseInt(artistIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid artistId: %w", err.Error())})
	}
	artist, err := s.db.GetArtist(c.Request().Context(), artistId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, artist)
}

func (s *Server) getAllAlbumsHandler(c echo.Context) error {
	albums, err := s.db.GetAllAlbums(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, albums)
}

func (s *Server) getAlbumHandler(c echo.Context) error {
	albumIdStr := c.Param("albumId")
	albumId, err := strconv.ParseInt(albumIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid albumId: %w", err.Error())})
	}
	album, err := s.db.GetAlbum(c.Request().Context(), albumId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, album)
}

func (s *Server) getAlbumsByArtistHandler(c echo.Context) error {
	artistIdStr := c.Param("artistId")
	artistId, err := strconv.ParseInt(artistIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid artistId: %w", err.Error())})
	}
	albums, err := s.db.GetAlbumsByArtist(c.Request().Context(), artistId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, albums)
}

func (s *Server) getTracksByAlbumHandler(c echo.Context) error {
	albumIdStr := c.Param("albumId")
	albumId, err := strconv.ParseInt(albumIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid albumId: %w", err.Error())})
	}
	tracks, err := s.db.GetTracksByAlbum(c.Request().Context(), albumId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, tracks)
}

func (s *Server) getPlaylistsHandler(c echo.Context) error {
	playlists, err := s.db.GetPlaylists(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, playlists)
}

func (s *Server) getPlaylistTracksHandler(c echo.Context) error {
	playlistIdStr := c.Param("playlistId")
	playlistId, err := strconv.ParseInt(playlistIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid playlistId: %w", err.Error())})
	}
	tracks, err := s.db.GetPlayListTracks(c.Request().Context(), playlistId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, tracks)
}
