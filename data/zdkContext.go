package data

import (
	"fmt"
	"os"

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
			log.Error("Failed to connect to database")
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
	ctx.ExecScript("clean.sql")
	ctx.AutoMigrate()
	ctx.ExecScript("defaults.sql")
	ctx.Log.Info("Done setting up database context")
}

func (ctx *ZdkContext) AutoMigrate() {
	ctx.DB.AutoMigrate(
		&models.Category{},
		&models.Equipment{},
		&models.Exercise{},
		&models.ExerciseDef{},
		&models.ScheduledTask{},
		&models.Setting{},
		&models.ShoppingList{},
		&models.ShoppingListItem{},
		&models.Task{},
		&models.TaskCategory{},
		&models.UnitType{},
		&models.UserSetting{},
		&models.Workout{},
	)
}

func (ctx *ZdkContext) ExecScript(fileName string) {
	input, err := os.ReadFile("./sql/" + fileName)
	if err != nil {
		ctx.Log.Error(err.Error())
		return
	}

	ctx.Log.Debug(fmt.Sprintf("Executing %s", fileName))

	ctx.DB.Exec(string(input))
}
