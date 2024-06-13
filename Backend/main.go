package main

import (
	"github.com/ekapamboedi/go-nbdTech/config"
	"github.com/ekapamboedi/go-nbdTech/model"
	"github.com/gin-gonic/gin"

	Employee "github.com/ekapamboedi/go-nbdTech/services/employee"
)

func main() {
	//Testing database connection here
	config.Init()
	model.Init()

	r := gin.Default()

	r.GET("/db-ping", func(ctx *gin.Context) {
		db, err := config.SupaCon()
		if err != nil {
			ctx.IndentedJSON(200, err.Error())
			defer db.Close()
		} else {
			response := map[string]string{}
			response["status"] = "Success"
			response["message"] = "Database successfully connected!"
			ctx.IndentedJSON(200, response)
		}
	})

	r.GET("", func(ctx *gin.Context) {
		response := map[string]string{}
		response["status"] = "Success"
		response["message"] = "Hello! You have been doing great things and it's matter!"
		ctx.IndentedJSON(200, response)
	})

	// needs to add middleware amd Authenticate

	// UserRoute := r.Group("/api/v1/user")
	// {
	// 	User.UserRoute(UserRoute)
	// }

	EmployeeRoute := r.Group("/api/v1/employee")
	{
		Employee.EmployeeRoute(EmployeeRoute)
	}
	r.Run(":" + config.SERVER_PORT)
}
