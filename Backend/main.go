package main

import (
	"github.com/ekapamboedi/go-project/config"
	"github.com/gin-gonic/gin"

	User "github.com/ekapamboedi/go-project/services/user"
)

func main() {
	config.Init()
	// config.InitOss()
	// model.Init()

	r := gin.Default()

	// Maximum Size for uploaded file
	// r.MaxMultipartMemory = 2 << 20

	r.GET("/db-ping", func(ctx *gin.Context) {
		db, err := config.DbCon()
		defer db.Close()
		if err != nil {
			panic("Database failed to connect: " + err.Error())
		}

		response := map[string]string{}
		response["status"] = "Success"
		response["message"] = "Database successfully connected!"
		ctx.IndentedJSON(200, response)
	})

	r.GET("", func(ctx *gin.Context) {
		response := map[string]string{}
		response["status"] = "Success"
		response["message"] = "Hello! You have been doing great things and it's matter!"
		ctx.IndentedJSON(200, response)
	})

	// needs to add middleware amd Authenticate
	UserRoute := r.Group("/api/v1/user")
	{
		User.UserRoute(UserRoute)
	}

	r.Run(":" + config.SERVER_PORT)
}
