package app

import (
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/adapter/postgres"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/controller/http_router"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/usecase"
)

func ProductDomain(d Dependences) {
	productUseCase := usecase.New(postgres.New(d.Postgres))

	http_router.ProductRoutes(d.RouterHTTP, productUseCase)
}
