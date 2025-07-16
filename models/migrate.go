package models

import (
	"foods-drinks-app/config"
)

func Migrate() error {
	return config.DB.AutoMigrate(
		&User{},
		&Product{},
		&Cart{},
		&Order{},
		&OrderItem{},
		&ProductReview{},
		&ProductSuggestion{},
	)
}
