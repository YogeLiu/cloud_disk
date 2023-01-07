package api

import (
	"net/http"

	"github.com/YogeLiu/CloudDisk/pkg/serializer"
	"github.com/YogeLiu/CloudDisk/server/user"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var u user.UserService
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ParamErr(err.Error(), err))
		return
	}
	err = u.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.DBErr(err.Error(), err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Response{})
}

func Login(ctx *gin.Context) {
	var u user.UserService
	err := ctx.ShouldBind(&u)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ParamErr(err.Error(), err))
		return
	}
	user, err := u.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.DBErr(err.Error(), err))
		return
	}
	token, err := u.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.Err(serializer.CodePasswordError, "password error", nil))
		return
	}
	ctx.JSON(http.StatusOK, serializer.DataResponse(gin.H{"token": token}))
}
