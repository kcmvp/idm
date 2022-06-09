package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kcmvp/idm/infra"
)

func main() {
	infra.SetupClient()
	engine := setupRouter()
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {

	router := gin.Default()

	accountGroup := router.Group("/account")
	{
		accountGroup.GET("")

	}

	applicationGroup := router.Group("/account")
	{
		applicationGroup.GET("")
	}

	roleGroup := router.Group("/account")
	{
		roleGroup.GET("")

	}

	funcGroup := router.Group("/account")
	{
		funcGroup.GET("")

	}

	return router
}
