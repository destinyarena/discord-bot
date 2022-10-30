package router

import "fmt"

type commandErrorWrapper struct {
	message string
	name    string
	err     error
}

func (e *commandErrorWrapper) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.name)
}

func (e *commandErrorWrapper) Unwrap() error {
	return e.err
}

func newCommandErrorWrapper(message, name string, err error) *commandErrorWrapper {
	return &commandErrorWrapper{message, name, err}
}

var (
	ErrCommandAlreadyRegistered = newCommandErrorWrapper("command already registered", "ErrCommandAlreadyRegistered", nil) // Command already registered
	ErrCommandNotFound          = newCommandErrorWrapper("command not found", "ErrCommandNotFound", nil)                   // Command not found
	ErrCommandHandlerNotFound   = newCommandErrorWrapper("command handler not found", "ErrCommandHandlerNotFound", nil)    // Command handler not found
)

func NewErrCommandAlreadyRegistered(name string) *commandErrorWrapper {
	err := ErrCommandAlreadyRegistered
	err.name = name
	return err
}

func NewErrCommandNotFound(name string) *commandErrorWrapper {
	err := ErrCommandNotFound
	err.name = name
	return err
}

func NewErrCommandHandlerNotFound(name string) *commandErrorWrapper {
	err := ErrCommandHandlerNotFound
	err.name = name
	return err
}

type componentErrorWrapper struct {
	err  string
	path string
}

func (e *componentErrorWrapper) Error() string {
	return fmt.Sprintf("%s: %s", e.err, e.path)
}

func newComponentErrorWrapper(err, path string) componentErrorWrapper {
	return componentErrorWrapper{err, path}
}

var (
	ErrComponentNotFound        = newComponentErrorWrapper("component not found", "ErrComponentNotFound")                // Component not found
	ErrComponentHandlerNotFound = newComponentErrorWrapper("component handler not found", "ErrComponentHandlerNotFound") // Component handler not found
)

func NewErrComponentNotFound(path string) *componentErrorWrapper {
	err := ErrComponentNotFound
	err.path = path
	return &err
}

func NewErrComponentHandlerNotFound(path string) *componentErrorWrapper {
	err := ErrComponentHandlerNotFound
	err.path = path
	return &err
}
