package router

type (
	ModuleHandlerFunc func(ctx *ModuleContext)

	ModuleInterface interface {
		Register(ctx *ModuleContext) error
		Commands() []*Command
		Components() []*Component
	}

	Module struct{}

	ModuleContext struct{}
)
