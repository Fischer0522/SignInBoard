package main

import (
	"SignInBoard/models"
	v1 "SignInBoard/routers/api/v1"
	"SignInBoard/setting"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, x-access-token")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
func main() {
	setting.Setup()
	models.Setup()

	r := gin.Default()
	r.Use(CORS())
	r.GET("/sign",v1.CurrentSign)

	r.POST("/sign",v1.CreateSignIn)

	r.PUT("/sign",v1.UpdateSign)

	r.GET("/trash",v1.GetAllTrash)

	r.DELETE("/trash/:id",v1.DeleteTrashById)

	r.POST("/trash",v1.GenerateTrash)


	r.Run()
	models.CloseDB()

}
