//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

// @author cold bin
// @date 2023/8/28
package main

import (
	"context"
	"design-pattern/di/foobarbaz"

	"github.com/google/wire"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.MegaSet)
	return foobarbaz.Baz{}, nil
}

func initializeBar(ctx context.Context) (foobarbaz.Bar, error) {
	wire.Build(foobarbaz.MegaSet)
	return foobarbaz.Bar{}, nil
}
