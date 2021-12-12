package fac

import (
	"testing"
)
//solid
//s 单一职责
//
//开闭 原则，不改老代码
//增加新功能
// litihuan 能用父类，能用自雷
// i 接口隔离
//d 以来导致
func TestNewFandian(t *testing.T) {
	NewFandian("d").Getfood()
	NewFandian("de").Getfood()
}