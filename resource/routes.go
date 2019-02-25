package resource

import (
	"github.com/efritz/chevron"
	"github.com/efritz/chevron/middleware"
	"github.com/efritz/nacelle"
)

var SetupRoutesFunc = chevron.RouteInitializerFunc(SetupRoutes)

func SetupRoutes(config nacelle.Config, router chevron.Router) error {
	router.AddMiddleware(middleware.NewLogging())
	router.AddMiddleware(middleware.NewRequestID())
	router.MustRegister("/", &QueryResource{})
	return nil
}
