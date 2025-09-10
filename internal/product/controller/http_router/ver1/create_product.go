package ver1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/error_type"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/response"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func (h *Handlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	input := dto.CreateProductInput{}

	invalidFields, err := validation.ValidateStructWithDecodeJSONBody(r.Body, &input)
	if err != nil {
		logger.Error(err, "validation.ValidateStructWithDecodeJSONBody")
		if invalidFields != nil {
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{Type: common_error.Validation, Details: invalidFields})
		} else {
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{Type: common_error.InvalidJSON})
		}
		return
	}

	output, err := h.uc.CreateProduct(ctx, input)
	if err != nil {
		logger.Error(err, "uc.CreateProduct")
		if errors.Is(err, error_type.CategoryDoesNotExist) {
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{
				Type:    error_type.CategoryDoesNotExist,
				Details: fmt.Sprintf("Category with code %s does not exist", input.CategoryCode),
			})
			return
		}
		response.Error(w, http.StatusInternalServerError, response.ErrorPayload{Type: common_error.InternalServer})
		return
	}

	response.Success(w, http.StatusCreated, output)
}
