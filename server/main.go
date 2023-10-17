package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/routes/general"
)

func main() {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.O.O.1"})
	if err != nil {
		fmt.Println("Error setting trusted proxies: ", err)
	}

	// Load routers
	general.RouterGeneral(r)

	err = r.Run(":3000")
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
