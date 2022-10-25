package router

type (
	CommandBuilder struct {
		command *Command
	}
)

func NewCommandBuilder(name string, description string) *CommandBuilder {
	return &CommandBuilder{
		command: &Command{
			Name:        name,
			Description: description,
			Commands:    NewCommandRouter(nil),
		},
	}
}

func (b *CommandBuilder) WithHandler(handler CommandHandlerFunc) *CommandBuilder {
	b.command.Handler = handler

	return b
}

func (b *CommandBuilder) WithOptions(options ...*CommandOption) *CommandBuilder {
	b.command.Options = options

	return b
}

func (b *CommandBuilder) WithSubCommands(commands ...*Command) *CommandBuilder {
	b.command.Commands.Register(commands...)

	return b
}

func (b *CommandBuilder) WithDefaultPermissions(permissions int64) *CommandBuilder {
	b.command.DefaultPermmissions = &permissions

	return b
}

func (b *CommandBuilder) WithDMPermissions(allowed bool) *CommandBuilder {
	b.command.DMPermmission = &allowed

	return b
}

func (b *CommandBuilder) Build() (*Command, error) {

	return b.command, nil
}

func (b *CommandBuilder) MustBuild() *Command {
	cmd, err := b.Build()
	if err != nil {
		panic(err)
	}

	return cmd
}
