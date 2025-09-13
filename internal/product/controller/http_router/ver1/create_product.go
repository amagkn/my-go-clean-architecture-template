package ver1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func (h *Handlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	input := dto.CreateProductInput{}

	invalidFields, err := validation.ValidateStructWithDecodeJSONBody(r.Body, &input)
	if err != nil {
		logger.Error(err, "validation.ValidateStructWithDecodeJSONBody")
		if invalidFields != nil {
			errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.Validation, Details: invalidFields})
		} else {
			errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.InvalidJSON})
		}
		return
	}

	output, err := h.uc.CreateProduct(ctx, input)
	if err != nil {
		logger.Error(err, "uc.CreateProduct")
		if errors.Is(err, entity.ErrCategoryDoesNotExist) {
			errorResponse(w, http.StatusBadRequest, errorPayload{
				Type:    entity.ErrCategoryDoesNotExist,
				Details: fmt.Sprintf("Category with code %s does not exist", input.CategoryCode),
			})
			return
		}
		errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		return
	}

	successResponse(w, http.StatusCreated, output)
}
