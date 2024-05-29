package User

import (
	"github.com/gin-gonic/gin"

	"github.com/ekapamboedi/go-project/services/user/response"
)

func UserRoute(r *gin.RouterGroup) {
	r.GET("/test", Test)
}

func Test(ctx *gin.Context) {
	var res response.Response
	res.Status = "Success"
	res.Message = "You are in user route"

	ctx.IndentedJSON(200, res)
}

func Register(ctx *gin.Context) {
	// var req request.RequestCreateUser
	// var res response.Response
	// var err error

	// err  = ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	res.Status = "Bad Request"
	// 	res.Message = "Error occured"
	// 	ctx.AbortWithStatusJSON(400, res)
	// 	go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
	// 	return
	// }

	// err = RegisterUser(req)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	res.Status = "Bad Request"
	// 	res.Message = err.Error()
	// 	ctx.AbortWithStatusJSON(400, res)
	// 	go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
	// 	return
	// }

	// res.Status = "Created"
	// res.Message = "Success"

	// ctx.IndentedJSON(201, res)
}
