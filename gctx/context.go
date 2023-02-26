package gctx

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Context struct {
	Ctx context.Context
}

func New(tx *gin.Context) {
	context.WithValue(tx, "ClientIP", "")
}
