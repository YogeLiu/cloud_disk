package api

import (
	"github.com/gin-gonic/gin"
)

func GetUploadSession(c *gin.Context) {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// var service explorer.CreateUploadSessionService
	// if err := c.ShouldBindJSON(&service); err == nil {
	// 	res := service.Create(ctx)
	// 	c.JSON(http.StatusOK, res)
	// } else {
	// 	c.JSON(http.StatusOK, ErrorResponse(err))
	// }
}
