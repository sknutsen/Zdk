package auth

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sknutsen/Zdk/config"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type UserContext struct {
	context.Context
}

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	log       *zap.Logger
	ZdkConfig *config.Config
}

func NewAuthenticator(log *zap.Logger, config *config.Config) *Authenticator {
	provider, err := oidc.NewProvider(
		context.Background(),
		config.AuthDomain+"/",
	)
	if err != nil {
		log.Sugar().Fatalf("Failed creating authenticator. Error: %s", err.Error())
	}

	conf := oauth2.Config{
		ClientID:     config.AuthClientId,
		ClientSecret: config.AuthClientSecret,
		RedirectURL:  config.AuthCallbackUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider:  provider,
		Config:    conf,
		log:       log,
		ZdkConfig: config,
	}
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
