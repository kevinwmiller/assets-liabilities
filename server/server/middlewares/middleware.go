package middlewares

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"assets-liabilities/logging"

	"github.com/gorilla/mux"
)

// AddContext returns a middleware function that will add the given context to the request
func AddContext(ctx context.Context) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Logging returns a middleware function that will log all incoming requests and where they came from
func Logging(ctx context.Context) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				logging.Logger(ctx).Debug(w, fmt.Sprintf("Requester ip: %q is not IP:port", r.RemoteAddr))
			}
			userIP := net.ParseIP(ip)

			logging.Logger(ctx).Infof("[%s]: %+v\n\n", userIP, r)
			next.ServeHTTP(w, r)
		})
	}
}
