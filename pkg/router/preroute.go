package router

type (
	PrerouteHandlerFunc func(ctx *PreRouteContext) error

	PreRouteContext struct{}

	PreRouteInterface interface{}
)
