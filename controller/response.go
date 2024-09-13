package controller

import (
	"FinancialAssistanceScheme/service"
	"github.com/kataras/iris/v12"
)

//const (
//	responseContextKey int = iota
//)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx iris.Context, data interface{}, status service.ServiceStatus) {
	resp := BaseResponse{
		Code:    status.Code,
		Message: status.Message,
		Data:    data,
	}

	if status != service.STATUS_OK {
		ctx.StatusCode(iris.StatusBadRequest)
	}

	ctx.JSON(resp)
}
