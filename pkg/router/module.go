package router

type (
	ModuleHandlerFunc func(ctx *ModuleContext)

	ModuleInterface interface {
		ModuleInit(ctx *ModuleContext) (*Module, error)
		Commands() []*Command
		Components() []*Component
		Modals() []*Modal
	}

	Module struct{}

	ModuleContext struct{}
)
