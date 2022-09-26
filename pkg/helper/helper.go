package helper

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type AllResponseData struct {
	OrderID      int            `json:"order_id"`
	OrderedAt    time.Time      `json:"ordered_at"`
	CustomerName string         `json:"customer_name"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}

func WriteJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Status, resp)
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err.Error(),
	}
}