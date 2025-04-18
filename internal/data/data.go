package data

import (
	"github.com/WH-5/friend-service/internal/conf"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gorm.io/driver/postgres"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFriendRepo)

// Data .
type Data struct {
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	err = MigrateDB(db)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		logHelper := log.NewHelper(logger)
		logHelper.Info("closing the data resources")

		sqlDB, err := db.DB()
		if err != nil {
			logHelper.Errorf("failed to get SQL DB: %v", err)
			return
		}

		// 关闭数据库连接并检查错误
		if err := sqlDB.Close(); err != nil {
			logHelper.Errorf("failed to close SQL DB: %v", err)
		}
	}
	return &Data{DB: db}, cleanup, nil
}
