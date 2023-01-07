package middle

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/YogeLiu/CloudDisk/dao"
	"github.com/YogeLiu/CloudDisk/pkg/secret"

	"github.com/gin-gonic/gin"
)

func CheckPerm(mode string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if mode == gin.TestMode {
			return
		}
		token := ctx.GetHeader("Authorization")
		if token == "" || token == "-1" {
			ctx.AbortWithError(http.StatusForbidden, errors.New("no authority"))
			return
		}
		id := secret.DeJwt(token)
		c, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		defer cancel()
		user, err := dao.GetUserByID(c, id)
		if err != nil {
			ctx.AbortWithError(http.StatusForbidden, errors.New("no authority"))
			return
		}
		if user.ID == 0 {
			ctx.AbortWithError(http.StatusForbidden, errors.New("no authority"))
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
