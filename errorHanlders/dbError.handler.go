package error_handlers

import "gorm.io/gorm"

func DBErrorHandler(result *gorm.DB) {
	if result.Error != nil {
		panic(result.Error)
	}
}