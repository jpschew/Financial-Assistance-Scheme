package model

import (
	"FinancialAssistanceScheme/middleware/db"
	"gorm.io/gorm/clause"
)

type Application struct {
	ID          uint64 `gorm:"column:id;type:int(11) unsigned;NOT NULL;primary_key;AUTO_INCREMENT" json:"id"`
	ApplicantID uint64 `gorm:"column:applicant_id;type:int(11) unsigned;NOT NULL" json:"applicant_id"`
	SchemeID    uint   `gorm:"column:scheme_id;type:int(11) unsigned;NOT NULL" json:"scheme_id"`
	Status      uint   `gorm:"column:status;type:int(11) unsigned;NOT NULL" json:"status"`
	CreateTime  uint64 `gorm:"column:create_time;type:int(11) unsigned;NOT NULL;autoCreateTime" json:"create_time"`
	UpdateTime  uint64 `gorm:"column:update_time;type:int(11) unsigned;NOT NULL;autoUpdateTime" json:"update_time"`
}

func (a *Application) Create() error {
	return db.GetDBConn().Clauses(clause.Insert{Modifier: "IGNORE"}).Create(a).Error
}

func (a *Application) Get() ([]*Application, error) {
	var results []*Application
	err := db.GetDBConn().Find(&results).Error
	return results, err
}

func (a *Application) GetByID() (*Application, error) {
	var result *Application
	err := db.GetDBConn().Model(a).Where(a, "id").First(&result).Error
	return result, err
}

func (a *Application) Update() error {
	return db.GetDBConn().Model(a).Where(a, "id").Update("status", a.Status).Error
}
