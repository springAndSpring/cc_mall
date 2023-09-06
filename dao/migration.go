package dao

import (
	"cc_mall/model"
	"fmt"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.Category{},
		&model.Cart{},
		&model.Carousel{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
		&model.User{},
	)
	if err != nil {
		fmt.Print("err", err)
	}
	return
}
