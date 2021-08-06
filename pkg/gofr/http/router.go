package http

import (
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"

	"github.com/vikash/gofr/pkg/gofr/logging"

	"github.com/vikash/gofr/pkg/gofr/http/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

func NewRouter() *Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	muxRouter.Use(
		middleware.Tracer,
		middleware.Logging(logging.NewLogger(logging.INFO)),
	)

	return &Router{
		Router: *muxRouter,
	}
}

func (rou *Router) Add(method, pattern string, handler http.Handler) {
	h := otelhttp.NewHandler(handler, fmt.Sprintf("%s %s", method, pattern))
	rou.Router.NewRoute().Methods(method).Path(pattern).Handler(h)
}
