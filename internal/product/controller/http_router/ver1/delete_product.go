package ver1

import (
	"context"
	"errors"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/internal/product/dto"
	"github.com/amagkn/my-go-clean-architecture-template/internal/product/entity"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/validation"
)

func (h *Handlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	input := dto.DeleteProductInput{}

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

	err = h.uc.DeleteProduct(ctx, input.ID)
	if err != nil {
		logger.Error(err, "uc.DeleteProduct")
		switch {
		case errors.Is(err, entity.ErrProductDoesNotExist):
			errorResponse(w, http.StatusBadRequest, errorPayload{
				Type:    entity.ErrProductDoesNotExist,
				Details: "Product with this ID does not exist",
			})
		default:
			errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		}
		return
	}

	successResponse(w, http.StatusOK, nil)
}
