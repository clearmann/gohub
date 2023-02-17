// 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

// 注册相关路由
func RegisterAPIRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			//判断手机号是否已经注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
