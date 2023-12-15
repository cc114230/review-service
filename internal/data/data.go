package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"review-service/internal/conf"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewReviewRepo, NewDB)

// Data .
type Data struct {
	// TODO wrapped database client
	// db *gorm.DB
	query *query.Query
	log   *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// 为GEN生成的query代码指定数据库连接对象
	query.SetDefault(db)
	return &Data{query: query.Q, log: log.NewHelper(logger)}, cleanup, nil
}

func NewDB(cfg *conf.Data) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.Database.GetSource()))
}
