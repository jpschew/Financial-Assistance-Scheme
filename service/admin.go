package service

import (
	"FinancialAssistanceScheme/middleware/jwt"
	"FinancialAssistanceScheme/middleware/redis"
	"FinancialAssistanceScheme/model"
	"FinancialAssistanceScheme/utils"
	"errors"
	"gorm.io/gorm"
	"log"
)

type Admin = model.Admin

type JWT struct {
	Token string `json:"token"`
}

const (
	RedisTimeoutMS = 15 * 60 * 1000
)

func CreateAdmin(name, username, password string) ServiceStatus {

	a := &Admin{
		Name:     name,
		Username: username,
		Password: password,
	}

	err := a.Create()
	if err != nil {
		return STATUS_DB_ERROR
	}

	return STATUS_OK
}

func Login(username, password string) (*JWT, ServiceStatus) {
	a := &Admin{
		Username: username,
	}

	admin, err := a.GetByUsername()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, STATUS_DB_ERROR
		}
		return nil, STATUS_NO_ADMIN_RECORD
	}

	if !utils.CheckPassword(admin.Password, password) {
		return nil, STATUS_PASSWORD_ERROR
	}

	key := "token_" + username
	token, err := redis.GetRedisManager().GetString(key)
	if err != nil {
		token, err = jwt.GenerateJWT(username, "admin")
		if err != nil {
			return nil, STATUS_GENERATE_TOKEN_ERROR
		}
		if err = redis.GetRedisManager().Set(key, token, RedisTimeoutMS); err != nil {
			log.Printf("set redis key %s failed %v\n", key, err)
		}
	}

	return &JWT{Token: token}, STATUS_OK

}

func Logout(username string) ServiceStatus {
	key := "token_" + username
	if err := redis.GetRedisManager().Del(key); err != nil {
		log.Printf("delete redis key %s failed %v\n", key, err)
		return STATUS_DEL_REDIS_TOKEN_ERROR
	}

	return STATUS_OK
}
