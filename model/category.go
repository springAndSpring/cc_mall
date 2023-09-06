package model

import "gorm.io/gorm"

// 商品分类
type Category struct {
	gorm.Model
	CategoryName string
}
