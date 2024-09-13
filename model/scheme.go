package model

import (
	"FinancialAssistanceScheme/middleware/db"
	"gorm.io/gorm"
)

//type CriteriaContent struct {
//	EmploymentStatus     int `json:"employment_status"` //",omitempty"`     // 0 - all, 1 - employed, 2 - unemployed
//	HasSchoolingChildren int `json:"has_children"`      //,omitempty"`   // 0 - all, 1 - has schooling children, 2 - without schooling children
//	MartialStatus        int `json:"martial_status"`    //,omitempty"` // 0 - all, 1 - single, 2 - married, 3 - divorced, 4 - widowed
//}

type Benefit struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Scheme struct {
	ID                uint   `gorm:"column:id;type:int(11) unsigned;NOT NULL;primary_key;AUTO_INCREMENT" json:"id"`
	Name              string `gorm:"column:name;type:text;NOT NULL" json:"name"`
	Description       string `gorm:"column:description;type:text;NOT NULL" json:"description"`
	EmployementStatus int    `gorm:"column:employment_status;type:int;NOT NULL" json:"employment_status"` // 0 - all, 1 - employed, 2 - unemployed
	MartialStatus     int    `gorm:"column:martial_status;type:int;NOT NULL" json:"martial_status"`       // 0 - all, 1 - single, 2 - married, 3 - divorced, 4 - widowed
	ChildrenStatus    int    `gorm:"column:children_status;type:int;NOT NULL" json:"children_status"`     // 0 - all, 1 - has schooling children, 2 - without schooling children
	//Criteria    *CriteriaContent `gorm:"column:criteria;type:JSON;NOT NULL;serializer:json" json:"criteria"`
	Benefits   []*Benefit `gorm:"column:benefits;type:JSON;NOT NULL;serializer:json" json:"benefits"` // need serializer to cast to json
	CreateTime uint64     `gorm:"column:create_time;type:int(11) unsigned;NOT NULL;autoCreateTime" json:"create_time"`
	UpdateTime uint64     `gorm:"column:update_time;type:int(11) unsigned;NOT NULL;autoUpdateTime" json:"update_time"`
}

func (s *Scheme) Create() error {
	return db.GetDBConn().Create(s).Error
}

func (s *Scheme) Get() ([]*Scheme, error) {
	var results []*Scheme
	err := db.GetDBConn().Find(&results).Error
	return results, err
}

func (s *Scheme) GetBySchemeID() (*Scheme, error) {
	var result *Scheme
	err := db.GetDBConn().Model(s).Where(s, "id").First(&result).Error
	return result, err
}

func (s *Scheme) GetEligibleScheme(employmentStatus, martialStatus, childrenStatus int) ([]*Scheme, error) {
	var results []*Scheme
	query := db.GetDBConn().Model(s).Scopes(FilterByEmploymentStatus(employmentStatus), FilterByMartialStatus(martialStatus)).Session(&gorm.Session{})
	//if employmentStatus != 0 {
	//	query = query.Scopes(FilterByEmploymentStatus(employmentStatus)).Session(&gorm.Session{})
	//}
	//if martialStatus != 0 {
	//	query = query.Scopes(FilterByMartialStatus(martialStatus)).Session(&gorm.Session{})
	//}
	if childrenStatus == 0 { // if no schooling children, filter out those that need schooling children
		query = query.Scopes(FilterByChildrenStatus).Session(&gorm.Session{})
	}

	err := query.Find(&results).Error
	return results, err
}

func FilterByChildrenStatus(db *gorm.DB) *gorm.DB {
	return db.Where("children_status = 0")
}

func FilterByEmploymentStatus(employmentStatus int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("employment_status = ? OR employment_status = 0", employmentStatus)
	}
}

func FilterByMartialStatus(martialStatus int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("martial_status = ? OR martial_status = 0", martialStatus)
	}
}

//func FilterByChildrenStatus(childrenStatus int) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		return db.Where("children_status = ?", childrenStatus)
//	}
//}
