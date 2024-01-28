package main

import (
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/config"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/handlers"
	"github.com/sknutsen/Zdk/middleware"
	"github.com/sknutsen/Zdk/routes"
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
			middleware.NewMiddleware,
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
