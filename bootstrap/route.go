package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleware(r)

	//注册api路由
	routes.RegisterAPIRouters(r)

	//配置404路由
	setup404Handle(r)
}
func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
}
func setup404Handle(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "页面返回404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "你访问的url不存在，请换个url重新访问",
			})
		}
	})
}
