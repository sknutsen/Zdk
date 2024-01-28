package middleware

import (
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/config"
)

type Middleware struct {
	Config         *config.Config
	SessionManager *auth.SessionManager
}

func NewMiddleware(config *config.Config, sm *auth.SessionManager) *Middleware {
	return &Middleware{Config: config, SessionManager: sm}
}
