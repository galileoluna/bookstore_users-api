package app

import (
	"github.com/galileoluna/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StarApplication() {

	mapUrls()
	logger.Info("about to start the aplicaction")
	router.Run(":8080")

}
