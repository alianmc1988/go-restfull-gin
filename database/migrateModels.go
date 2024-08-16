package database

import (
	user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"
	"gorm.io/gorm"
)

var Models  = []interface {}{
	&user_models.User{},
}

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(Models...)

}