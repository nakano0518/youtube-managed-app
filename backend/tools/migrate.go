package main

import (
	"youtube-managed-app/backend/databases"
	"youtube-managed-app/backend/models"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect() //返却されるdbは*gorm.DB
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	//*gorm.DBの有するAutoMigrateは引数にモデル(のポインタ)を与えることで、モデルの定義を読み取りマイグレーションを実行する
	//テーブル定義
	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
