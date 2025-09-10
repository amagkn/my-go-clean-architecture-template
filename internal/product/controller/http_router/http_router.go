package http_router

import (
	ver1 "github.com/amagkn/my-go-clean-architecture-template/internal/product/controller/http_router/v1"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/usecase"
	"github.com/go-chi/chi/v5"
)

func ProductRouter(r *chi.Mux, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	r.Route("/api/v1/product", func(r chi.Router) {
		r.Get("/all", v1.GetProducts)
		r.Post("/create", v1.CreateProduct)
		r.Patch("/update", v1.UpdateProduct)
		r.Delete("/delete", v1.DeleteProduct)
	})
}
