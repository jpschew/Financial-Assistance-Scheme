package controller

import (
	"FinancialAssistanceScheme/service"
	"github.com/kataras/iris/v12"
)

type Application struct {
	Name string
}

type EmptyData struct {
}

func GetAllApplications(ctx iris.Context) {

	applications, status := service.GetAllApplications()
	SendResponse(ctx, applications, status)

}

func CreateApplication(ctx iris.Context) {

	var application *ApplicationInfo

	// Read JSON body into the application struct
	if err := ctx.ReadJSON(&application); (err != nil && !iris.IsErrPath(err)) || application.SchemeID == 0 || application.NRIC == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.CreateApplication(application.NRIC, application.SchemeID)
	SendResponse(ctx, &EmptyData{}, status)
}

func UpdateApplication(ctx iris.Context) {

	var application *UpdateStatus

	// Read JSON body into the application struct
	if err := ctx.ReadJSON(&application); (err != nil && !iris.IsErrPath(err)) || application.ID == 0 || application.Status > 3 {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.UpdateApplication(application.ID, application.Status)
	SendResponse(ctx, &EmptyData{}, status)
}
