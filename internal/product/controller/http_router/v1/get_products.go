package v1

import (
	"context"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/response"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func (h *Handlers) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	input := dto.GetProductsInput{
		CategoryCode: r.URL.Query().Get("categoryCode"),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")
		response.Error(w, http.StatusBadRequest, response.ErrorPayload{Type: common_error.Validation, Details: invalidFields})
		return
	}

	output, err := h.uc.GetProducts(ctx, input)
	if err != nil {
		logger.Error(err, "uc.GetProducts")
		response.Error(w, http.StatusInternalServerError, response.ErrorPayload{Type: common_error.InternalServer})
		return

	}
	response.Success(w, http.StatusOK, output)
}
