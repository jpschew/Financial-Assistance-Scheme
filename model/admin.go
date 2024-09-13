package model

import (
	"FinancialAssistanceScheme/middleware/db"
	"FinancialAssistanceScheme/utils"
	"gorm.io/gorm"
)

type Admin struct {
	ID         uint64 `gorm:"column:id;type:int(11) unsigned;NOT NULL;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `gorm:"column:name;type:text;NOT NULL" json:"name"`
	Username   string `gorm:"column:username;type:text;NOT NULL" json:"username"`
	Password   string `gorm:"column:password;type:text;NOT NULL" json:"password"`
	CreateTime uint64 `gorm:"column:create_time;type:int(11) unsigned;NOT NULL;autoCreateTime" json:"create_time"`
	UpdateTime uint64 `gorm:"column:update_time;type:int(11) unsigned;NOT NULL;autoUpdateTime" json:"update_time"`
}

func (a *Admin) BeforeSave(db *gorm.DB) error {
	// If the password is not already hashed, hash it
	if len(a.Password) < 60 { // bcrypt hashed password length is 60
		hashedPassword, err := utils.HashPassword(a.Password)
		if err != nil {
			return err
		}
		a.Password = hashedPassword
	}
	return nil
}

func (a *Admin) Create() error {
	return db.GetDBConn().Create(a).Error
}

func (a *Admin) GetByUsername() (*Admin, error) {
	var result *Admin
	err := db.GetDBConn().Model(a).Where(a, "username").First(&result).Error
	return result, err
}
