package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchRelatedVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		videoId := c.Param("id")

		call := yts.Search.
			List([]string{"id", "snippet"}).
			RelatedToVideoId(videoId).
			Type("video"). //RelatedToVideoIdをしてする際この記述必要(要ドキュメント参照)
			MaxResults(3)
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling YouTube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
