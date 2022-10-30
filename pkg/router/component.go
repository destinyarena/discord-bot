package router

type (
	ComponentHandlerFunc func(ctx *ComponentContext) error

	ComponentContext struct {
		*Context
		Path   string
		Params map[string]string
		Values []string
	}

	Component struct {
		Path    string
		Params  map[string]string
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

func (r *ComponentRouter) Get(path string) (*Component, error) {
	v, params := r.components.Search(path)
	if v == nil {
		return nil, NewErrComponentNotFound(path)
	}

	handler := v.(ComponentHandlerFunc)
	if handler == nil {
		return nil, NewErrComponentHandlerNotFound(path)
	}

	return &Component{
		Path:    path,
		Params:  params,
		Handler: handler,
	}, nil
}
