package middlewares

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) { //firebase.goでc.Set("firebase", auth)したので使用時それを引数に渡す
	//リクエストヘッダーからトークンの取出し(付属した文字列Bearerは除く)
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	//auth.VerifyIDTokenにトークンを与えることで検証できる(firebase)
	//jwtTokenはUIDを持つ構造体//このUIDをお気に入り追加・削除機能で使用する
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)

	return jwtToken, err
}

//お気に入り追加・削除機能などのログインした状態でなければ利用できないAPIに対して使用
func FirebaseGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := verifyFirebaseIDToken(c, authClient)

			c.Set("auth", jwtToken) //検証後JWTトークン(UIDを持つ構造体データ)をcontextにSet

			if err != nil {
				return c.JSON(fasthttp.StatusUnauthorized, "Not Authenticated")
			}

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

//動画を既にお気に入り登録しているかどうかに認証情報を利用するが、ログインしていなくても利用できるAPIに対して使用
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := verifyFirebaseIDToken(c, authClient)

			c.Set("auth", jwtToken)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
