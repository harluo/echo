package echo

import (
	"github.com/harluo/echo/internal/builder"
	"github.com/harluo/echo/internal/core"
	"github.com/harluo/echo/internal/kernel"
)

type Route = core.Route

func NewRoute[Q any, S any](target core.Target, handler kernel.Handler[Q, S]) *builder.Route[Q, S] {
	return builder.NewRoute(target, handler)
}
