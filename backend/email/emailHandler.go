package email

import (
	"github.com/gin-gonic/gin"
)

func SetupEmailRoutes(router *gin.Engine) {
	router.POST("/email/send", SendEmail)
}