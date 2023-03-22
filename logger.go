package gin_utils

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const ctxKey = "ctxKey"

func Get(ctx context.Context) *zerolog.Logger {
	return ctx.Value(ctxKey).(*zerolog.Logger)
}

func WithLogger(ctx *gin.Context, l zerolog.Logger) {
	ctx.Set(ctxKey, &l)
}
