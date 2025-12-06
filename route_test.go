package echo_test

import (
	"testing"

	"github.com/harluo/echo"
	"github.com/harluo/echo/internal/kernel"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	server := new(echo.Server)
	assert.NotNil(t, echo.NewRoute(server, func(ctx *kernel.Context, req *int) (rsp string, err error) {
		return "Hello World", nil
	}))
}
