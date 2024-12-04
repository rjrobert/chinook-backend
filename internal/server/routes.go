package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rjrobert/chinook-music/backend/internal/repository/database"
	"golang.org/x/time/rate"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	e.GET("/health", s.healthHandler)

	e.GET("/artists", s.getArtistsHandler)
	e.GET("/artists/:artistId", s.getArtistHandler)
	e.GET("/artists/:artistId/albums", s.getAlbumsByArtistHandler)

	e.GET("/albums", s.getAllAlbumsHandler)
	e.GET("/albums/:albumId", s.getAlbumHandler)
	e.GET("/albums/:albumId/tracks", s.getTracksByAlbumHandler)

	e.GET("/playlists", s.getPlaylistsHandler)
	e.GET("/playlists/:playlistId", s.getPlaylistTracksHandler)

	e.GET("/genres", s.getGenresHandler)
	e.GET("/mediatypes", s.getMediaTypesHandler)

	//TODO: Should be protected routes
	e.GET("/customers", s.getCustomersHandler)
	e.POST("/customers", s.createCustomerHandler)
	e.GET("/customers/:customerId/invoices", s.getCustomerInvoicesHandler)
	e.GET("/customers/:customerId/invoices/:invoiceId", s.getInvoiceDetailsHandler)

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

func (s *Server) getGenresHandler(c echo.Context) error {
	genres, err := s.db.GetGenres(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, genres)
}

func (s *Server) getMediaTypesHandler(c echo.Context) error {
	mediaTypes, err := s.db.GetMediaTypes(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, mediaTypes)
}

func (s *Server) getCustomersHandler(c echo.Context) error {
	customers, err := s.db.GetCustomers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, customers)
}

func (s *Server) getCustomerInvoicesHandler(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.ParseInt(customerIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid customerId: %w", err.Error())})
	}
	invoices, err := s.db.GetCustomerInvoices(c.Request().Context(), customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, invoices)
}

func (s *Server) getInvoiceDetailsHandler(c echo.Context) error {
	invoiceIdStr := c.Param("invoiceId")
	invoiceId, err := strconv.ParseInt(invoiceIdStr, 0, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Invalid invoiceId: %w", err.Error())})
	}
	invoices, err := s.db.GetInvoiceDetails(c.Request().Context(), invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, invoices)
}

func (s *Server) createCustomerHandler(c echo.Context) error {
	req := new(database.CreateCustomerParams)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid customer"})
	}
	newCustomer, err := s.db.CreateCustomer(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, newCustomer)
}
