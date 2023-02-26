package gsvc

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	g  *gin.RouterGroup
	ir gin.IRouter
}

func (s *Server) NewRouter(name string) *Router {
	// 私有化的路由请求
	g := s.e.Group(name)
	return &Router{g: g}
}

func (r *Router) Use(handlerFunc ...gin.HandlerFunc) *Router {
	r.g.Use(handlerFunc...)
	return r
}

func (r *Router) handler(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.g.Handle(httpMethod, relativePath, handlers...)
}

func (r *Router) Post(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.POST(api, handlerFunc...)
	return r
}

func (r *Router) Get(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.GET(api, handlerFunc...)
	return r
}

func (r *Router) Delete(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.DELETE(api, handlerFunc...)
	return r
}

func (r *Router) Put(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.PUT(api, handlerFunc...)
	return r
}

func (r *Router) Patch(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.PATCH(api, handlerFunc...)
	return r
}

func (r *Router) Options(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.OPTIONS(api, handlerFunc...)
	return r
}

func (r *Router) Head(api string, handlerFunc ...gin.HandlerFunc) *Router {
	r.g.HEAD(api, handlerFunc...)
	return r
}

func (r *Router) StaticFile(s1 string, s2 string) *Router {
	r.g.StaticFile(s1, s2)
	return r
}
