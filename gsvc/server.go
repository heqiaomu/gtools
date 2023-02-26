package gsvc

import (
	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"time"
)

type Server struct {
	cfg *Config
	e   *gin.Engine
	op  Options
	r   *gin.RouterGroup
}

func NewServer(cfg *Config, ops ...Options) *Server {
	s := &Server{cfg: cfg}
	for _, op := range ops {
		op(s)
	}
	if s.cfg.AppDebug {
		s.e = gin.Default()
		pprof.Register(s.e)
	} else {
		s.e = newEngine()
	}
	return s
}

func (s *Server) Use(handlerFunc ...gin.HandlerFunc) *Server {
	s.e.Use(handlerFunc...)
	return s
}

func (s *Server) Start() {
	svc := initServer(s.cfg.Server.API.Port, s.e)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	tls := s.cfg.Server.TLS
	if tls.Enabled {
		svc.ListenAndServeTLS(tls.Cert, tls.Key)
	} else {
		svc.ListenAndServe()
	}
}

type server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
