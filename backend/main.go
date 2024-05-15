package main

import (
	"chouseisan/handler"
	"chouseisan/repository"
	"chouseisan/schedule"
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	router := gin.Default()
	// repository.InitDB()
	// http.HandleFunc("/", setCookies)
	// http.HandleFunc("/cookie", showCookie)
	// solve the CORS block problem
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.GET("/refreshToken", cookie.RefreshCookieHandler)
	// router.GET("/setCookie", cookie.SetCookieHandler)
	// router.GET("/cookie", getCookieHandler)

	// router.GET("/add-text", func(c *gin.Context) {
	// 	// Use the context's String method to add text to the response
	// 	c.String(http.StatusOK, "Hello, this is some text added to the response!")
	// })

	schedule.SetupScheduleRoutes(router)
	repo := repository.NewRepository(repository.DB)

	event_handler := handler.NewEventHandler(repo)

	// router.GET("/eventBasic/:uuid", event_handler.EventBasicHandler)

	router.POST("/createEvent", event_handler.CreateEventHandler)
	// event, err := repository.GetEventByTitle()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Event found: %v\n", event)
	// events.SetupEventsRoutes(router)

	handler.SetupEventRoutes(router, event_handler)

	c := cron.New()
	_, err := c.AddFunc("@every 1m", event_handler.CheckDueDatesAndSendEmails)
	if err != nil {
		log.Fatal("Error setting up the scheduler:", err)
	}
	c.Start()
	defer c.Stop()
	
	router.Run("0.0.0.0:8080")
}
