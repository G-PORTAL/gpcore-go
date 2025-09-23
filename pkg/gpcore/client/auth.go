package client

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

type AuthProviderOption interface {
	GetToken(context.Context) (string, error)
	Impersonate(token *gocloak.JWT)
}

type AuthOption struct {
	Provider *AuthProviderOption
}

func (a *AuthOption) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if (a.Provider == nil) || (*a.Provider == nil) {
		return nil, fmt.Errorf("no provider set")
	}
	token, err := (*a.Provider).GetToken(ctx)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", token),
	}, nil
}

func (a *AuthOption) RequireTransportSecurity() bool {
	return true
}
