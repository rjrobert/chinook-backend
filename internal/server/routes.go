package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	e.GET("/albums", s.getAlbumsHandler)
	e.GET("/albums/:albumId", s.getAlbumHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getAlbumsHandler(c echo.Context) error {
	albums, err := s.db.GetAlbums(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, albums)
}

func (s *Server) getAlbumHandler(c echo.Context) error {
	albumId := c.Param("albumId")
	// return c.JSON(http.StatusOK, map[string]string{"albumId": albumId})
	album, err := s.db.GetAlbum(context.WithValue(c.Request().Context(), "albumId", albumId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, album)
}
