package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/logger"
	"github.com/go-my-admin/server/routes/connections"
	"github.com/go-my-admin/server/routes/general"
	"github.com/go-my-admin/server/routes/interactDatabases"
	"github.com/go-my-admin/server/utils/bootstrap"
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

	// Load middlewares
	r.Use(gin.Recovery()) // Recover from panics and return 500 in case of panic

	// Load routes
	general.Router(r)
	connections.Router(r)
	interactDatabases.Router(r)

	err = r.Run(":3000")
	if err != nil {
		logger.Error("Error running server: ", err)
		return
	}
}
