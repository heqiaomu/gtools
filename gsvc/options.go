package gsvc

import "github.com/gin-gonic/gin"

type Options func(s *Server)

func WithMiddleware(middle ...gin.HandlerFunc) Options {
	return func(s *Server) {
		s.Use(middle...)
	}
}
