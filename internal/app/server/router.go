package server

import (
	"os"
	"time"

	"kbrprime-be/internal/app/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
)

func Router(opt handler.HandlerOption) *gin.Engine {
	healthyHandler := handler.HealthyCheckHandler{
		HandlerOption: opt,
	}

	showHandler := handler.ShowHandler{
		HandlerOption: opt,
	}

	setMode := cast.ToBool(os.Getenv("DEBUG_MODE"))
	if setMode {
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	//routes
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders:           []string{"Origin", "Content-Type", "Authorization", "userid", "REQUEST-ID", "X-SIGNATURE", "Referer", "User-Agent"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 120 * time.Second,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	r.Use(gin.Recovery())

	//Maximum memory limit for Multipart forms
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("healthy-check", healthyHandler.HealthyCheck)
	}
	showGroup := r.Group("/api/v1/show")
	{
		showGroup.GET("/all", showHandler.GetAll)
		showGroup.GET("/latest-news", showHandler.LatestNews)
		showGroup.GET("/latest-episodes", showHandler.LatestEpisodes)
		showGroup.GET("/list-news", showHandler.ListNews)
		showGroup.GET("/list-episodes", showHandler.ListEpisodes)
	}

	return r
}
