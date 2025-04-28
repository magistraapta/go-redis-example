package seed

import (
	"redis-caching/model"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	users := []model.User{
		{ID: 1, Username: "john", Email: "john@email.com", Password: "john123"},
		{ID: 2, Username: "mike", Email: "mike@email.com", Password: "mike123"},
	}

	for _, user := range users {
		var existingUser model.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&user)
			}
		}
	}
}
