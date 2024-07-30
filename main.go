package main

import (
	"net/http"

	template "crud/src"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, "public/*.html")
	e.GET("/Hello", func(e echo.Context) error {
		res := map[string]interface{}{
			"Name":  "Muthoimo",
			"Phone": "071111111",
			"Email": "Muthoimo12@gmail.com",
		}

		return e.Render(http.StatusOK, "index", res)
	})
	e.Logger.Fatal(e.Start(":8083"))
}
