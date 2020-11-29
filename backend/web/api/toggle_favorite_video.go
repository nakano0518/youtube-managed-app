package api

import (
	"youtube-managed-app/backend/middlewares"
	"youtube-managed-app/backend/models"

	"firebase.google.com/go/auth"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type ToggleFavoriteVideoResponse struct {
	VideoId    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() { //userテーブルにトークンから取得したUIDを保持するレコードが存在しない場合
			//新規ユーザーをusersテーブルに格納する
			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		if dbs.DB.Table("favorites").
			Where(models.Favorite{UserId: user.ID, VideoId: videoId}).First(&favorite).RecordNotFound() { //favoritesテーブルに条件に合致するレコードが存在しない場合=>追加

			favorite = models.Favorite{UserId: user.ID, VideoId: videoId}
			dbs.DB.Create(&favorite)
			isFavorite = true
		} else { //favoritesテーブルに条件に合致するレコードが存在する場合=>削除
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:    videoId,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
