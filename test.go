package main

type App struct {
	Name     string
	Database string
}

type Router[T any] struct {
	Data T
}

func (r *Router[any]) Get[T any]() T {
	return r.Data
}

func main() {

	app := App{
		Name:     "Discord Bot",
		Database: "MongoDB",
	}

	r := Router[App]{Data: app}

	println(r.Get().Name)

}
