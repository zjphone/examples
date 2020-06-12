package middleware

import (
	"github.com/casbin/casbin"
	"github.com/kataras/iris/context"
	"net/http"
)

func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}

func (c *Casbin) Wrapper() func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		if !c.Check(r) {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("403 Forbidden"))
			return
		}
		router(w, r)
	}
}

func (c *Casbin) ServeHTTP(ctx context.Context) {
	if !c.Check(ctx.Request()) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbiden
		ctx.StopExecution()
		return
	}
	ctx.Next()
}

type Casbin struct {
	enforcer *casbin.Enforcer
}

func (c *Casbin) Check(r *http.Request) bool {
	username := Username(r)
	method := r.Method
	path := r.URL.Path
	b := c.enforcer.Enforce(username, path, method)
	return b
}

func Username(r *http.Request) string {
	//username, _, _ := r.BasicAuth()//这玩意我用不上，把它注释掉
	return "abc123" //直接返回用户名，看看测试效果
}
