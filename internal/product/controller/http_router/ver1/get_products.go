package ver1

import (
	"context"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
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
		errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.Validation, Details: invalidFields})
		return
	}

	output, err := h.uc.GetProducts(ctx, input)
	if err != nil {
		logger.Error(err, "uc.GetProducts")
		errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		return

	}
	successResponse(w, http.StatusOK, output)
}
