package controller

import (
	"FinancialAssistanceScheme/service"
	"github.com/kataras/iris/v12"
)

type Scheme struct {
	Name        string
	Description string
}

func GetAllSchemes(ctx iris.Context) {

	schemes, status := service.GetAllSchemes()
	SendResponse(ctx, schemes, status)

}

func CreateScheme(ctx iris.Context) {
	var scheme *SchemeInfo

	// Read JSON body into the scheme struct
	if err := ctx.ReadJSON(&scheme); (err != nil && !iris.IsErrPath(err)) || scheme.Name == "" || scheme.Description == "" || len(scheme.Benefits) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	if scheme.EmploymentStatus == 0 && scheme.ChildrenStatus == 0 && scheme.MartialStatus == 0 { // check for empty criteria
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.CreateScheme(scheme.Name, scheme.Description, scheme.EmploymentStatus, scheme.MartialStatus, scheme.ChildrenStatus, scheme.Benefits)
	SendResponse(ctx, &EmptyData{}, status)

}

func GetEligibleScheme(ctx iris.Context) {

	// Read URL Param
	applicantID, err := ctx.URLParamInt64("applicant_id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_APPLICANT_ID_NOT_INTEGER)
		return
	}

	schemes, status := service.GetEligibleSchemes(uint64(applicantID))
	SendResponse(ctx, schemes, status)

}
