package web

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

// 这样写可以分组可以防止main 太大 分散路由
func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	server.POST("users/signup", u.SignUp)
	server.POST("users/login", u.Login)
	server.POST("users/edit", u.Edit)
	server.GET("users/profile", u.Profile)
}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

}
func (u *UserHandler) Login(ctx *gin.Context) {

}
func (u *UserHandler) Edit(ctx *gin.Context) {

}
func (u *UserHandler) Profile(ctx *gin.Context) {

}
