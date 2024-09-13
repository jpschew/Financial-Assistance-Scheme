package controller

import (
	"FinancialAssistanceScheme/service"
	"github.com/kataras/iris/v12"
)

func GetAllApplicants(ctx iris.Context) {

	applicants, status := service.GetAllApplicants()
	SendResponse(ctx, applicants, status)

}

func CreateApplicant(ctx iris.Context) {

	var applicant *ApplicantInfo

	// Read JSON body into the applicant struct
	if err := ctx.ReadJSON(&applicant); (err != nil && !iris.IsErrPath(err)) || applicant.FirstName == "" || applicant.LastName == "" || applicant.DateOfBirth == "" || applicant.NRIC == "" || applicant.EmploymentStatus == 0 || applicant.Sex > 1 {
		ctx.StatusCode(iris.StatusBadRequest)
		SendResponse(ctx, &EmptyData{}, service.STATUS_PARAMS_EMPTY)
		return
	}

	status := service.CreateApplicant(applicant.FirstName, applicant.LastName, applicant.NRIC, applicant.EmploymentStatus, applicant.MartialStatus, applicant.Sex, applicant.DateOfBirth, applicant.Household)
	SendResponse(ctx, &EmptyData{}, status)

}
