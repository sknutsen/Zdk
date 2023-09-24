package auth

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/data"
)

type SessionManager struct {
	ZdkCtx       *data.ZdkContext
	SessionStore *session.Store
	Config       *config.Config
}

func NewSessionManager(ctx *data.ZdkContext, config *config.Config) *SessionManager {
	return &SessionManager{ZdkCtx: ctx, Config: config, SessionStore: session.New()}
}
