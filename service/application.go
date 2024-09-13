package service

import (
	"FinancialAssistanceScheme/model"
	"errors"
	"gorm.io/gorm"
)

type Application = model.Application

func GetAllApplications() ([]*Application, ServiceStatus) {

	a := &Application{}
	allApplicants, err := a.Get()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, STATUS_DB_ERROR
		}
		return nil, STATUS_NO_APPLICATION_RECORD
	}

	return allApplicants, STATUS_OK
}

func CreateApplication(nric string, schemeID uint) ServiceStatus {

	if !validateNRIC(nric) {
		return STATUS_INVALID_NRIC
	}

	a := &Applicant{
		NRIC: nric,
	}

	applicant, err := a.GetByNRIC()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return STATUS_DB_ERROR
		}
		return STATUS_NO_APPLICANT_RECORD
	}

	s := &Scheme{
		ID: schemeID,
	}

	scheme, err := s.GetBySchemeID()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return STATUS_DB_ERROR
		}
		return STATUS_NO_SCHEME_RECORD
	}

	app := &Application{
		ApplicantID: applicant.ID,
		SchemeID:    scheme.ID,
	}

	err = app.Create()
	if err != nil {
		return STATUS_DB_ERROR
	}

	return STATUS_OK
}

func UpdateApplication(applicationID uint64, status uint) ServiceStatus {
	app := &Application{
		ID: applicationID,
	}

	application, err := app.GetByID()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return STATUS_DB_ERROR
		}
		return STATUS_NO_APPLICATION_RECORD
	}

	if application.Status == status {
		if status == APPLICATION_PENDING {
			return STATUS_APPLICATION_STILL_PENDING
		} else if status == APPLICATION_APPROVED {
			return STATUS_APPLICATION_ALREADY_APPROVED
		} else if status == APPLICATION_REJECTED {
			return STATUS_APPLICATION_ALREADY_REJECTED
		} else if status == APPLICATION_WITHDRAWN {
			return STATUS_APPLICATION_ALREADY_WITHDRAWN
		}
	}

	application.Status = status

	err = application.Update()
	if err != nil {
		return STATUS_UPDATE_APPLICATION_STATUS_FAILED
	}

	return STATUS_OK

}
