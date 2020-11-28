package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		call := yts.Videos.
			List([]string{"id", "snippet"}).
			Chart("mostPopular").
			MaxResults(3)
		pageToken := c.QueryParam("pageToken") //NuxtからのリクエストにクエリパラメータでpageTokenが設定されている場合
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
