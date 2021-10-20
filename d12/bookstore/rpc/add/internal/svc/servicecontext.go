package svc

import (
	"bookstore/rpc/add/internal/config"
	"dailytest/daily_test/d12/bookstore/rpc/model"
)

type ServiceContext struct {
	c config.Config
	Model model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
		Model:model.NewBookModel(sqlx.NewMysql(c.DataSource),c.Cache),

}
}