package svc

import (
	"bookstore/api/internal/config"
	"dailytest/daily_test/d12/bookstore/rpc/add/adder"
	"dailytest/daily_test/d12/bookstore/rpc/check/checker"
)

type ServiceContext struct {
	Config config.Config
	Adder adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Adder:adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
