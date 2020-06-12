package main

import (
	"github.com/kataras/iris"

	cm "examples/casbin-test/middleware"
	"github.com/casbin/casbin"
)

var e = casbin.NewEnforcer("E:\\workspace\\go\\src\\examples\\casbin-test\\rbac_model.conf", "E:\\workspace\\go\\src\\examples\\casbin-test\\rbac_policy.csv")

func newApp() *iris.Application {
	casbinMiddleware := cm.New(e)
	app := iris.New()
	app.WrapRouter(casbinMiddleware.Wrapper())
	// 如果不想使用中间件，可以通过下面方法进行判断
	/*
	   if b,err:=e.Enforce("abc123","/user","Get");b {
	       fmt.Println("成功")
	   } else {
	       fmt.Println("失败")
	   }
	*/
	app.Get("/user", hi)
	app.Post("/user", hi)
	app.Put("/test", hi)

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func hi(ctx iris.Context) {
	ctx.Writef("当你看到这个，说明通过了权限验证")
}
