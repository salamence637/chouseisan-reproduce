package schedule

import (
	"github.com/gin-gonic/gin"
)

func SetupScheduleRoutes(router *gin.Engine) {
	router.GET("/chouseisan/schedule", GetMembersAvailability)
	router.POST("chouseisan/schedule", PostMembersAvailability)
}
