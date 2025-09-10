package main

import (
	"context"

	"github.com/amagkn/my-go-clean-architecture-template/config"
	"github.com/amagkn/my-go-clean-architecture-template/internal/app"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(c.Logger)
	validation.Init()

	err = app.Run(ctx, c)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
