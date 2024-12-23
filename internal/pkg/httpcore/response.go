package httpcore

import (
	"net/http"

	"github.com/chhz0/goiam/internal/pkg/errorscore"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code       int
	HttpStatus int
	Message    string
	Reference  string
}

func WriteResponse(ctx *gin.Context, err error, data any) {
	if err != nil {
		coder := errorscore.ParseErrorToCoder(err)

		// todo any的类型转换
		var ref string
		var ok bool
		if coder.Any() != nil {
			ref, ok = coder.Any().(string)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse{
					Code:       coder.Code(),
					HttpStatus: coder.HttpStauts(),
					Message:    coder.Message(),
				})
				return
			}
		}

		ctx.JSON(coder.HttpStauts(), ErrorResponse{
			Code:       coder.Code(),
			HttpStatus: coder.HttpStauts(),
			Message:    coder.Message(),
			Reference:  ref,
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}
