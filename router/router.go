package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gof/app/api"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/hello", api.Hello.Index)
		group.GET("/in", api.Hello.Insert)
		group.GET("/t", api.Hello.Find)
		group.GET("/redis", api.Hello.Redis)
		group.GET("/addr", api.Hello.Addr)
	})
}
