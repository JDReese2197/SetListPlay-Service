package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jdreese/setlistplay/api/handlers"
	"net/http"
)

var e *echo.Echo

func Init() {
	e = echo.New()

	//	Middleware
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAccept, echo.HeaderOrigin},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	//	Routes
	//	SetlistFM
	e.GET("/setlistfm/event/search", handlers.HandleEventSearch)
	e.GET("/setlistfm/setlist/search", handlers.HandleSetlistSearch)
	//Spotify
	e.GET("/spotify/login", handlers.HandleSpotifyLogin)

	//	Test / Example
	e.GET("/test", handlers.SetlistFmTest)
}

func Run() {
	e.Logger.Fatal(e.Start(":3000"))
}
