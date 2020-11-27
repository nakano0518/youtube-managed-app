package routes

import (
	"youtube-managed-app/backend/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos()) //endpoint("/api/popular")
	}
}
