package controller

import "FinancialAssistanceScheme/model"

type HouseholdContent = model.HouseholdContent

// type CriteriaContent = model.CriteriaContent
type Benefit = model.Benefit

type ApplicantInfo struct {
	FirstName        string              `json:"first_name"`
	LastName         string              `json:"last_name"`
	NRIC             string              `json:"nric"`
	EmploymentStatus int                 `json:"employment_status"`
	MartialStatus    int                 `json:"martial_status"`
	Sex              int                 `json:"sex"`
	DateOfBirth      string              `json:"date_of_birth"`
	Household        []*HouseholdContent `json:"household"`
}

type ApplicationInfo struct {
	SchemeID uint   `json:"scheme_id"`
	NRIC     string `json:"nric"`
}

type UpdateStatus struct {
	ID     uint64 `json:"application_id"`
	Status uint   `json:"status"`
}

type SchemeInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	//Criteria    *CriteriaContent `json:"criteria"`
	EmploymentStatus int        `json:"employment_status"`
	MartialStatus    int        `json:"martial_status"`
	ChildrenStatus   int        `json:"children_status"`
	Benefits         []*Benefit `json:"benefits"`
}

type AdminInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutInfo struct {
	Username string `json:"username"`
}
