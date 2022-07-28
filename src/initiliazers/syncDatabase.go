package initiliazers

import "github.com/0xlilnas/shopapp/src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
