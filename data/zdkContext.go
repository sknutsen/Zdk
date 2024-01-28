package data

import (
	"github.com/sknutsen/Zdk/config"
	"github.com/sknutsen/Zdk/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IZdkContext interface {
	SetupContext()
	AutoMigrate(dst ...interface{}) error
}

type ZdkContext struct {
	DB     *gorm.DB
	Config *config.Config
	Log    *zap.Logger
}

var zCtx *ZdkContext

func NewZdkContext(log *zap.Logger, config *config.Config) (*ZdkContext, error) {
	if zCtx == nil {
		ctx, err := openDb(config.DbType, config.DbHost, config.DbName, config.DbUser, config.DbPass, config.DbPort)
		if err != nil {
			log.Debug("Failed to connect to database")
			return nil, err
		}

		zCtx = &ZdkContext{DB: ctx, Config: config, Log: log}

		if config.DbUpdate {
			zCtx.SetupContext()
		}
	}

	return zCtx, nil
}

func (ctx *ZdkContext) SetupContext() {
	ctx.Log.Info("Setting up database context")
	ctx.DB.AutoMigrate(
		&models.Category{},
		&models.ScheduledTask{},
		&models.Setting{},
		&models.ShoppingList{},
		&models.ShoppingListItem{},
		&models.Task{},
		&models.TaskCategory{},
		&models.UserSetting{},
		&models.Workout{},
	)
}
