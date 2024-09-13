package service

import (
	"FinancialAssistanceScheme/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Scheme = model.Scheme
type Benefit = model.Benefit

type CriteriaContent struct {
	EmploymentStatus int `json:"employment_status,omitempty"` // 0 - all, 1 - employed, 2 - unemployed
	ChildrenStatus   int `json:"children_status,omitempty"`   // 0 - all, 1 - has schooling children, 2 - without schooling children
	MartialStatus    int `json:"martial_status,omitempty"`    // 0 - all, 1 - single, 2 - married, 3 - divorced, 4 - widowed
}

type SchemeResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Criteria    *CriteriaContent `json:"criteria"`
	Benefits    []*Benefit       `json:"benefits"`
}

func GetAllSchemes() ([]*SchemeResponse, ServiceStatus) {

	s := &Scheme{}
	allSchemes, err := s.Get()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, STATUS_DB_ERROR
		}
		return nil, STATUS_NO_SCHEME_RECORD
	}

	schemes := makeSchemeResponse(allSchemes)

	return schemes, STATUS_OK

}

func CreateScheme(name, description string, employmentStatus, martialStatus, childrenStatus int, benefits []*Benefit) ServiceStatus {
	s := &Scheme{
		Name:              name,
		Description:       description,
		EmployementStatus: employmentStatus,
		MartialStatus:     martialStatus,
		ChildrenStatus:    childrenStatus,
		Benefits:          benefits,
	}

	err := s.Create()
	if err != nil {
		return STATUS_DB_ERROR
	}

	return STATUS_OK
}

func GetEligibleSchemes(applicantID uint64) ([]*SchemeResponse, ServiceStatus) {
	a := &Applicant{
		ID: applicantID,
	}

	var schemes []*SchemeResponse

	applicant, err := a.GetByID()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return schemes, STATUS_DB_ERROR
		}
		return schemes, STATUS_NO_APPLICANT_RECORD
	}

	c := &CriteriaContent{
		MartialStatus:    applicant.MartialStatus,
		EmploymentStatus: applicant.EmploymentStatus,
	}

	for _, person := range applicant.Household {
		if strings.ToLower(person.Relation) == "son" || strings.ToLower(person.Relation) == "daughter" {
			birthYear, _ := strconv.Atoi(strings.Split(person.DateOfBirth, "-")[0])
			currentYear := time.Now().Year()
			if currentYear-birthYear <= 16 && currentYear-birthYear >= 7 { // only for primary to secondary school
				c.ChildrenStatus = 1
				break
			}
		}
	}

	s := &Scheme{}

	eligibleSchemes, err := s.GetEligibleScheme(c.EmploymentStatus, c.MartialStatus, c.ChildrenStatus)
	if err != nil {
		//if !errors.Is(err, gorm.ErrRecordNotFound) {
		//	return schemes, STATUS_DB_ERROR
		//}
		return schemes, STATUS_DB_ERROR
	}

	schemes = makeSchemeResponse(eligibleSchemes)

	return schemes, STATUS_OK

}

func makeSchemeResponse(schemes []*Scheme) []*SchemeResponse {
	resp := make([]*SchemeResponse, len(schemes))

	for i, scheme := range schemes {
		resp[i] = &SchemeResponse{
			ID:          scheme.ID,
			Name:        scheme.Name,
			Description: scheme.Description,
			Criteria: &CriteriaContent{
				EmploymentStatus: scheme.EmployementStatus,
				MartialStatus:    scheme.MartialStatus,
				ChildrenStatus:   scheme.ChildrenStatus,
			},
			Benefits: scheme.Benefits,
		}
	}

	return resp
}
