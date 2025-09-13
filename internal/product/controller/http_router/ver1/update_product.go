package ver1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/common_error"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/response"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func (h *Handlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	input := dto.UpdateProductInput{}

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

	err = h.uc.UpdateProduct(ctx, input)
	if err != nil {
		logger.Error(err, "uc.UpdateProduct")
		switch {
		case errors.Is(err, entity.ErrCategoryDoesNotExist):
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{
				Type:    entity.ErrCategoryDoesNotExist,
				Details: fmt.Sprintf("Category with code %s does not exist", *input.CategoryCode),
			})
		case errors.Is(err, entity.ErrProductDoesNotExist):
			response.Error(w, http.StatusBadRequest, response.ErrorPayload{
				Type:    entity.ErrProductDoesNotExist,
				Details: "Product with this ID does not exist",
			})
		default:
			response.Error(w, http.StatusInternalServerError, response.ErrorPayload{Type: common_error.InternalServer})
		}
		return
	}

	response.Success(w, http.StatusOK, nil)
}
