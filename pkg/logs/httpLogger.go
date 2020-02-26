package logs

import (
	"io"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
)

// GetHTTPLoggerMiddleware get http logger middleware
func GetHTTPLoggerMiddleware(out io.Writer, info map[string]string) alice.Chain {
	logger := NewLogger(out, info)

	middlewares := alice.New()
	// Install the logger handler with logger
	middlewares = middlewares.Append(hlog.NewHandler(*logger.Logger))

	// Install some provided extra handler to set some request's context fields.
	middlewares = middlewares.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("receive request body")
	}))
	middlewares = middlewares.Append(hlog.RemoteAddrHandler("ip"))
	middlewares = middlewares.Append(hlog.UserAgentHandler("user_agent"))
	middlewares = middlewares.Append(hlog.RefererHandler("referer"))
	middlewares = middlewares.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	return middlewares
}
