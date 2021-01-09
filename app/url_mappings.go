package app

import (
	"github.com/tjhooper1/bookstore_users-api/controllers"
)
func mapUrls(){
	router.GET("/ping", controllers.Ping)
}