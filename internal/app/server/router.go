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

	authHandler := handler.AuthHandler{
		HandlerOption: opt,
	}

	userHandler := handler.UserHandler{
		HandlerOption: opt,
	}

	showHandler := handler.ShowHandler{
		HandlerOption: opt,
	}

	categoriesHandler := handler.CategoriesHandler{
		HandlerOption: opt,
	}

	listenHandler := handler.ListenHandler{
		HandlerOption: opt,
	}

	likeHandler := handler.LikeHandler{
		HandlerOption: opt,
	}
	playListHandler := handler.PlayListHandler{
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
		apiGroup.GET("/healthy-check", healthyHandler.HealthyCheck)

		apiGroup.GET("/berita-terbaru", showHandler.LatestNews)
		apiGroup.GET("/episode-terbaru", showHandler.LatestEpisodes)
		apiGroup.GET("/list-berita", showHandler.ListNews)
		apiGroup.GET("/list-episode", showHandler.ListEpisodes)
		apiGroup.POST("/register", userHandler.AddNewUser)
		apiGroup.POST("/record", listenHandler.RecordData)
		apiGroup.POST("/like", likeHandler.LikeEpisode)
		apiGroup.GET("/teratas-minggu-ini", showHandler.TopThree)
		apiGroup.GET("/sorotan", showHandler.Sorotan)
	}
	authAPIGroupUnsecured := r.Group("/api/v1/auth")
	{
		authAPIGroupUnsecured.POST("/login", authHandler.Login)
	}
	userGroup := r.Group("/api/v1/user", opt.AuthMiddleware.AuthorizeUser())
	{
		userGroup.POST("/register", userHandler.AddNewUser)
		userGroup.GET("/all", userHandler.GetAllUser)
		userGroup.PUT("/update", userHandler.UpdateUser)
		userGroup.GET("/:id", userHandler.GetDetailUser)
		userGroup.GET("/logout", authHandler.RevokeToken)
	}

	playListGroup := r.Group("/api/v1/play-list", opt.AuthMiddleware.AuthorizeUser())
	{
		playListGroup.GET("/", playListHandler.GetPlayList)
		playListGroup.POST("/add", playListHandler.AddPlayList)
		playListGroup.DELETE("delete/:id", playListHandler.DeletePlayList)
	}

	showGroup := r.Group("/api/v1/show")
	{
		showGroup.GET("/all", showHandler.GetAll)
	}

	categroriesGroup := r.Group("/api/v1/categories")
	{
		categroriesGroup.GET("/list", categoriesHandler.GetAll)
	}

	return r
}
