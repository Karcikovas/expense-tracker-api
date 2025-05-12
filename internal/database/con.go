package database

import (
	"expense-tracker-api/internal/config"
	"expense-tracker-api/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBCon(cfg config.Config, logger logger.Service) *gorm.DB {
	dbCon, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Database.ConnectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.LogError(err.Error())
	}

	return dbCon
}
