package data

import (
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/models"
	"gorm.io/gorm"
)

type IZdkContext interface {
	SetupContext()
	AutoMigrate(dst ...interface{}) error
}

type ZdkContext struct {
	DB     *gorm.DB
	Config *config.Config
}

var zCtx *ZdkContext

func NewZdkContext(config *config.Config) (*ZdkContext, error) {
	if zCtx == nil {
		ctx, err := openDb(config.DbType, config.DbHost, config.DbName, config.DbUser, config.DbPass, config.DbPort)
		if err != nil {
			println("Failed to connect to database")
			return nil, err
		}

		zCtx = &ZdkContext{DB: ctx, Config: config}
	}

	return zCtx, nil
}

func (ctx *ZdkContext) SetupContext() {
	ctx.AutoMigrate(&models.ShoppingList{})
}

func (ctx *ZdkContext) AutoMigrate(dst ...interface{}) error {
	err := ctx.DB.AutoMigrate(dst)

	return err
}
