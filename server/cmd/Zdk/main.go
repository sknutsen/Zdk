package main

import (
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/data"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/routes"
	"github.com/sknutsen/Zdk/internal/storage"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			data.NewZdkContext,
			storage.NewShoppingListsStorage,
			storage.NewShoppingListItemsStorage,
			handlers.NewShoppingListsHandler,
			handlers.NewShoppingListItemsHandler,
		),
		fx.Invoke(routes.NewRouter),
	).Run()
}
