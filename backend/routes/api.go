package routes

import (
	"youtube-managed-app/backend/middlewares"
	"youtube-managed-app/backend/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos()) //endpoint("/api/popular")
		g.GET("/video/:id", api.GetVideo(), middlewares.FirebaseAuth())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}
	//ログインしたユーザーのみ使用できる機能
	fg := g.Group("/favorite", middlewares.FirebaseGuard()) //'/api'のグループを更にグルーピング// '/api/favorite'配下にアクセスする場合middlewares.FirebaseGuard()を経由している必要がある
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}
