package helper

import (
	"assignment2/pkg/params"
	"net/http"

	"github.com/gin-gonic/gin"
)
func WriteJsonRespnse(ctx *gin.Context, resp *params.Response) {
	ctx.JSON(resp.Status, resp)
}

func SuccessCreateResponse(payload interface{}, message string) *params.Response {
	return &params.Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessResponse(payload interface{}, message string) *params.Response {
	return &params.Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *params.Response {
	return &params.Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
	}
}