package router

type (
	ModuleHandlerFunc func(ctx *ModuleContext)

	ModuleInterface interface {
		Register(ctx *ModuleContext) error
		Commands() []*Command
		Components() []*Component
		Modals() []*Modal
	}

	Module struct{}

	ModuleContext struct{}
)
