package main

import (
	"github.com/ekapamboedi/go-nbdTech/config"
	"github.com/gin-gonic/gin"

	Employee "github.com/ekapamboedi/go-nbdTech/services/employee"
)

func main() {
	// config.PgCon()
	// config.InitOss()
	// model.Init()
	config.SupaCon()

	r := gin.Default()

	// // r.GET("/db-ping", func(ctx *gin.Context) {
	// // 	db, err := config.DbCon()
	// // 	defer db.Close()
	// // 	if err != nil {
	// // 		panic("Database failed to connect: " + err.Error())
	// // 	}

	// 	response := map[string]string{}
	// 	response["status"] = "Success"
	// 	response["message"] = "Database successfully connected!"
	// 	ctx.IndentedJSON(200, response)
	// })

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
