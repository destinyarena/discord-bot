package router

import "fmt"

type (
	ModalExistsError struct {
		ModalName string
	}

	ModalInvalidArgumentError struct {
		ModalName string
		Err       error
	}

	ModalNotFoundError struct {
		ModalName string
	}

	ComponentExistsError struct {
		ComponentName string
	}

	ComponentInvalidArgumentError struct {
		ComponentName string
		Err           error
	}

	ComponentNotFoundError struct {
		ComponentName string
	}

	CommandAlreadyRegisteredError struct {
		CommandName string
	}

	CommandNotFoundError struct {
		CommandName string
	}

	CommandExitError struct {
		CommandName string
		Err         error
	}

	ErrSubCommandGroupAlreadyExists struct {
		GroupName string
	}
)

func (e *ErrSubCommandGroupAlreadyExists) Error() string {
	return fmt.Sprintf("subcommand group %s already exists", e.GroupName)
}

func (e *ModalExistsError) Error() string {
	return fmt.Sprintf("Modal already exists: %s\n", e.ModalName)
}

func (e *ModalInvalidArgumentError) Error() string {
	return fmt.Sprintf("Modal %s has invalid arguments: %s\n", e.ModalName, e.Err.Error())
}

func (e *ModalNotFoundError) Error() string {
	return fmt.Sprintf("Modal not found: %s\n", e.ModalName)
}

func (e *ComponentExistsError) Error() string {
	return fmt.Sprintf("Component already exists: %s\n", e.ComponentName)
}

func (e *ComponentInvalidArgumentError) Error() string {
	return fmt.Sprintf("Component %s has invalid arguments: %s\n", e.ComponentName, e.Err.Error())
}

func (e *ComponentNotFoundError) Error() string {
	return fmt.Sprintf("Component not found: %s\n", e.ComponentName)
}

func (e *CommandAlreadyRegisteredError) Error() string {
	return fmt.Sprintf("Command already exists: %s\n", e.CommandName)
}

func (e *CommandNotFoundError) Error() string {
	return fmt.Sprintf("Command not found: %s\n", e.CommandName)
}

func (e *CommandExitError) Error() string {
	return fmt.Sprintf("Command %s exited with error: %s\n", e.CommandName, e.Err.Error())
}
