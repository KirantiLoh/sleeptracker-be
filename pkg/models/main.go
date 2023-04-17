package models

import "github.com/kirantiloh/dashtrack-be/pkg/config"

func init() {
	config.Connect()
	db := config.GetDB()

	db.AutoMigrate()
}
