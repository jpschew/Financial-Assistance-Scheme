package controller

import (
	"FinancialAssistanceScheme/service"
	"github.com/kataras/iris/v12"
)

func CreateAdmin(ctx iris.Context) {

	var admin *AdminInfo

	// Read JSON body into the admin struct
	if err := ctx.ReadJSON(&admin); (err != nil && !iris.IsErrPath(err)) || admin.Name == "" || admin.Username == "" || admin.Password == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.CreateAdmin(admin.Name, admin.Username, admin.Password)
	SendResponse(ctx, &EmptyData{}, status)

}

func Login(ctx iris.Context) {
	var login *LoginInfo

	// Read JSON body into the login struct
	if err := ctx.ReadJSON(&login); (err != nil && !iris.IsErrPath(err)) || login.Username == "" || login.Password == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	token, status := service.Login(login.Username, login.Password)
	SendResponse(ctx, token, status)

}

func Logout(ctx iris.Context) {
	var logout *LogoutInfo

	// Read JSON body into the logout struct
	if err := ctx.ReadJSON(&logout); (err != nil && !iris.IsErrPath(err)) || logout.Username == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.Logout(logout.Username)
	SendResponse(ctx, &EmptyData{}, status)
}
