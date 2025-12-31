package internal

import (
	"github.com/harluo/echo/internal/kernel"
	"github.com/labstack/echo/v4"
)

type Next struct {
	next echo.HandlerFunc
}

func NewNext(next echo.HandlerFunc) *Next {
	return &Next{
		next: next,
	}
}

func (n *Next) Execute(ctx *kernel.Context) error {
	return n.next(ctx.Echo())
}
