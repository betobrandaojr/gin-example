package main

import (
	"log"

	_ "gin-gonic-example/docs"
	"gin-gonic-example/src/configuration/util"
	"gin-gonic-example/src/controller/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	util.LoadEnv()

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup)

	port := ":" + util.Env.PORT
	log.Printf("Server starting on port %s", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
