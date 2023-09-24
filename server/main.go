package main

import (
	"github.com/sknutsen/Zdk/internal/auth"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/data"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/routes"
	"github.com/sknutsen/Zdk/internal/storage"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewExample,
			config.NewConfig,
			data.NewZdkContext,
			auth.NewAuthenticator,
			auth.NewSessionManager,
			storage.NewShoppingListsStorage,
			storage.NewShoppingListItemsStorage,
			storage.NewCategoriesStorage,
			storage.NewScheduledTasksStorage,
			storage.NewSettingsStorage,
			storage.NewTaskCategoriesStorage,
			storage.NewTasksStorage,
			storage.NewUserSettingsStorage,
			storage.NewUsersStorage,
			handlers.NewCallbackHandler,
			handlers.NewLoginHandler,
			handlers.NewLogoutHandler,
			handlers.NewShoppingListsHandler,
			handlers.NewShoppingListItemsHandler,
			handlers.NewCategoriesHandler,
			handlers.NewScheduledTasksHandler,
			handlers.NewSettingsHandler,
			handlers.NewTaskCategoriesHandler,
			handlers.NewTasksHandler,
			handlers.NewUserHandler,
			handlers.NewUserSettingsHandler,
			handlers.NewUsersHandler,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(routes.NewRouter),
	).Run()
}
