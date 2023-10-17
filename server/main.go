package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/logger"
	"github.com/go-my-admin/server/routes/general"
	"github.com/go-my-admin/server/utils/bootstrap"
)

var InternalDb database.DBConnection

func main() {
	// load env variables
	err := bootstrap.LoadEnv()
	if err != nil {
		logger.Error("Error loading env variables: ", err)
		return
	}

	// Init connection to internal database
	if bootstrap.BootInternalDb(&InternalDb) != nil {
		logger.Error("Error bootstrapping internal database: ", err)
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

	// Load routers
	general.RouterGeneral(r)

	err = r.Run(":3000")
	if err != nil {
		logger.Error("Error running server: ", err)
		return
	}
}
