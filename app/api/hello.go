package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gof/app/service"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {

	res, _ := service.Test.Test()
	g.Dump(res)
}

func (*helloApi) Insert(r *ghttp.Request) {

	service.Insert()
}

func (*helloApi) Find(r *ghttp.Request) {

	service.Find()
}

func (*helloApi) Redis(r *ghttp.Request) {

	service.Redis()
}

func (*helloApi) Addr(r *ghttp.Request) {

	service.Addr()
}
