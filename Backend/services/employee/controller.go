package Employee

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ekapamboedi/go-nbdTech/services/employee/request"
	"github.com/ekapamboedi/go-nbdTech/services/employee/response"
)

func EmployeeRoute(r *gin.RouterGroup) {
	r.GET("/test", Test)
}

func Test(ctx *gin.Context) {
	var res response.Response
	res.Status = "Success"
	res.Message = "You are in Employee route"

	ctx.IndentedJSON(200, res)
}

func Create(ctx *gin.Context) {
	var req request.RequestCreateEmployee
	var res response.Response
	var err error

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		// go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	err = CreateOne(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		return
	}

	res.Status = "Created"
	res.Message = "Success"

	ctx.IndentedJSON(201, res)
}
