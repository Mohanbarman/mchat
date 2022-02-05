package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"mchat.com/api/config"
	"mchat.com/api/db"
	"mchat.com/api/modules/ws/connection"
	"mchat.com/api/router"
	"mchat.com/api/validation"
)

func main() {
	r := gin.Default()
	config := config.New()
	db := db.New(config.Database)

	if db == nil {
		panic("Failed to connect to database")
	}

	ws := connection.NewStore()

	validation.UseJsonKeyTagName()
	router.SetupRoutes(r, config, db, ws)

	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}
