package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/docs"
	"github.com/go-my-admin/server/logger"
	"github.com/go-my-admin/server/routes/auth"
	"github.com/go-my-admin/server/routes/connections"
	"github.com/go-my-admin/server/routes/general"
	"github.com/go-my-admin/server/routes/interactDatabases"
	"github.com/go-my-admin/server/utils/bootstrap"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// load env variables
	err := bootstrap.LoadEnv()
	if err != nil {
		logger.Error("Error loading env variables: ", err)
		return
	}

	// Init connection to internal database
	if bootstrap.BootInternalDb() != nil {
		logger.Error("Error bootstrapping internal database: ", err)
		return
	}

	// Init connection to internal cache
	if bootstrap.BootInternalCache() != nil {
		logger.Error("Error bootstrapping internal cache: ", err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Error("Error setting trusted proxies: ", err)
		return
	}

	docs.SwaggerInfo.BasePath = "/"

	// Load middlewares
	r.Use(gin.Recovery()) // Recover from panics and return 500 in case of panic
	r.Use(cors.Default()) // Enable CORS for all routes

	// Load routes
	general.Router(r)
	auth.Router(r)
	connections.Router(r)
	interactDatabases.Router(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Info("Server running on port 3000")
	logger.Info("Swagger running on http://localhost:3000/swagger/index.html")
	err = r.Run(":3000")
	if err != nil {
		logger.Error("Error running server: ", err)
		return
	}
}
