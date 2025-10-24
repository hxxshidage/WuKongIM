package api

import (
	"errors"
	"github.com/WuKongIM/WuKongIM/internal/options"
	"github.com/WuKongIM/WuKongIM/pkg/wkhttp"
	"go.uber.org/zap"
	"net/http"
)

func apiCallMw(logFunc func(string, ...zap.Field)) func(c *wkhttp.Context) {
	return func(c *wkhttp.Context) {
		valFromInternal := c.GetHeader("k-from-internal")
		internalKey := options.G.HttpInternalKey
		if internalKey != "" && valFromInternal != internalKey {
			err := errors.New("forbidden: not from trusted server")

			logFunc("access forbidden", zap.Error(err))

			c.ResponseErrorWithStatus(http.StatusForbidden, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
