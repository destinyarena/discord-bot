package router

type (
	PrerouteHandlerFunc func(ctx *PreRouteContext) error

	PreRouteContext struct{}

	PreRouteInterface interface{}

	PreRoute struct {
		Handler   PrerouteHandlerFunc
		NextRoute *PreRoute
	}
)
