package router

/*
\/button\/name\/(?<number>[A-z0-9]+)\/(?<fsadfasdf>[A-z0-9]+)
*/

type (
	ComponentHandlerFunc func(ctx *ComponentContext)

	ComponentContext struct{}

	Component struct {
		Path    string
		Handler ComponentHandlerFunc
	}

	ComponentRouter struct {
		components *Node
	}
)

func NewComponentRouter(components []*Component) *ComponentRouter {
	router := &ComponentRouter{
		components: NewTree(),
	}

	for _, component := range components {
		router.Register(component)
	}

	return router
}

func (r *ComponentRouter) Register(components ...*Component) error {
	for _, component := range components {
		r.components.Insert(component.Path, component.Handler)
	}

	return nil
}

func (r *ComponentRouter) Get(path string) ComponentHandlerFunc {
	v := r.components.Search(path)
	if v == nil {
		return nil
	}

	return v.(ComponentHandlerFunc)
}
