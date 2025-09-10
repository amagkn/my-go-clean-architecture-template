package v1

import (
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/usecase"
)

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{uc: uc}
}
