package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HandleEventSearch
//
//	Accepts - Artist name,  date, year, venue, city, state, country
//	Returns - List of events (concerts) matching search criteria
func HandleEventSearch(c echo.Context) error {
	artist := c.QueryParam("artist")
	date := c.QueryParam("date")
	year := c.QueryParam("year")
	venue := c.QueryParam("venue")
	city := c.QueryParam("city")
	state := c.QueryParam("state")
	country := c.QueryParam("country")

	return c.String(http.StatusOK, fmt.Sprintf(`Artist: %s, Date: %s, Year: %s, Venue: %s, Location: %s, %s, %s`, artist, date, year, venue, city, state, country))
}
