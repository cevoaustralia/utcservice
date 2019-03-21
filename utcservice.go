package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/araddon/dateparse"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", parse)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func parse(c echo.Context) error {
	input := c.QueryParam("date")
	tz := c.QueryParam("tz")

	t, err := parseDate(input, tz)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, t.String()+"\n")
}

func parseDate(input, tz string) (time.Time, error) {
	if "" == input {
		return time.Unix(0, 0), echo.NewHTTPError(http.StatusBadRequest, "Missing query param: date")
	}

	if "" != tz {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return time.Unix(0, 0), echo.NewHTTPError(http.StatusBadRequest, "Invalid timezone: "+tz)
		}
		time.Local = loc
	} else {
		time.Local = time.UTC
	}

	t, err := dateparse.ParseLocal(input)
	if err != nil {
		return time.Unix(0, 0), echo.NewHTTPError(http.StatusInternalServerError, "Could not parse date")
	}

	return t, nil
}
