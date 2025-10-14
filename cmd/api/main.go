package main

import (
	_ "gin-gonic-example/docs"
	"gin-gonic-example/src/configuration/util"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Gonic Example API
// @version 1.0
// @description Simple API example with Gin
// @BasePath /
func main() {
	util.LoadEnv()

	r := gin.Default()

	// @Summary Hello World
	// @Description Get hello world message
	// @Tags hello
	// @Produce json
	// @Success 200 {object} map[string]string
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger host: http://localhost:" + util.Env.PORT + "/swagger/index.html")
	r.Run(":" + util.Env.PORT)
}
