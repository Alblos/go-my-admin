package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/general"
	"github.com/go-my-admin/server/utils/bootstrap"
)

func main() {

	// load env variables
	err := bootstrap.LoadEnv()
	if err != nil {
		fmt.Println("Error loading env variables: ", err)
		return
	}

	err = database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	r := gin.Default()
	err = r.SetTrustedProxies([]string{"*"})
	if err != nil {
		fmt.Println("Error setting trusted proxies: ", err)
		return
	}

	// Load middlewares
	r.Use(gin.Recovery()) // Recover from panics and return 500 in case of panic

	// Load routers
	general.RouterGeneral(r)

	err = r.Run(":3000")
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
