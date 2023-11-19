package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HandlePathParam(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func HandleQueryParams(c echo.Context) error {
	name := c.QueryParam("name")
	team := c.QueryParam("team")
	return c.String(http.StatusOK, "Name: "+name+", team: "+team)
}
