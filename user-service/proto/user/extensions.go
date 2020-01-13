package laracom_user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User)BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id",id.String())
	// create database `laracom-user-db` charset=utf8mb4 collate utf8mb4_general_ci;
}
