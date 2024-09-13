package model

import (
	"FinancialAssistanceScheme/middleware/db"
)

type HouseholdContent struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NRIC        string `json:"nric"`
	Sex         int    `json:"sex"`
	Relation    string `json:"relation"`
	DateOfBirth string `json:"date_of_birth"`
}

type Applicant struct {
	ID               uint64              `gorm:"column:id;type:int(11) unsigned;NOT NULL;primary_key;AUTO_INCREMENT" json:"id"`
	FirstName        string              `gorm:"column:first_name;type:text;NOT NULL" json:"first_name"`
	LastName         string              `gorm:"column:last_name;type:text;NOT NULL" json:"last_name"`
	NRIC             string              `gorm:"column:nric;type:text;NOT NULL" json:"nric"`
	EmploymentStatus int                 `gorm:"column:employment_status;type:int;NOT NULL" json:"employment_status"` // 1 - employed, 2 - unemployed
	MartialStatus    int                 `gorm:"column:martial_status;type:int;NOT NULL" json:"martial_status"`       // 1 - single, 2 - married, 3 - divorced, 4 - widowed
	Sex              int                 `gorm:"column:sex;type:int;NOT NULL" json:"sex"`
	DateOfBirth      string              `gorm:"column:date_of_birth;type:text;NOT NULL" json:"date_of_birth"`
	Household        []*HouseholdContent `gorm:"column:household;type:JSON;NOT NULL;serializer:json" json:"household"` // need serializer to cast to json
	CreateTime       uint64              `gorm:"column:create_time;type:int(11) unsigned;NOT NULL;autoCreateTime" json:"create_time"`
	UpdateTime       uint64              `gorm:"column:update_time;type:int(11) unsigned;NOT NULL;autoUpdateTime" json:"update_time"`
}

func (a *Applicant) Create() error {
	return db.GetDBConn().Create(a).Error
}

func (a *Applicant) Get() ([]*Applicant, error) {
	var results []*Applicant
	err := db.GetDBConn().Find(&results).Error
	return results, err
}

func (a *Applicant) GetByNRIC() (*Applicant, error) {
	var result *Applicant
	err := db.GetDBConn().Model(a).Where(a, "nric").First(&result).Error
	return result, err
}

func (a *Applicant) GetByID() (*Applicant, error) {
	var result *Applicant
	err := db.GetDBConn().Model(a).Where(a, "id").First(&result).Error
	return result, err
}
