package handler

import (
	"net/http"
	"url-shortenr/database"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func ResolverURL(c echo.Context) error {

	url := c.Param("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.String(http.StatusNotFound, "short not found in database")
	} else if err != nil {
		return c.String(http.StatusInternalServerError, "can't connect to database")
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(301, value)
}
