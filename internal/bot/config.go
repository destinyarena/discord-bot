package bot

type (
	Config struct {
		// Bot Stuff
		Prefix string `yaml:"prefix"`
		Token  string `yaml:"token"`

		// Guild Stuff
		GuildID      string `yaml:"guild_id"`
		LogsID       string `yaml:"logs_id"`
		JoinRoleID   string `yaml:"join_role"`
		BannedRoleID string `yaml:"banned_role"`

		// Permissions
		Owners     []string `yaml:"owners"`
		AdminRoles []string `yaml:"admin_roles"`
		ModRoles   []string `yaml:"mod_roles"`

		// Reaction roles stuff
		RegistrationRoleID string          `yaml:"registration_role"`
		Hubs               []*Hub          `yaml:"hubs"`
		ReactionRoles      []*ReactionRole `yaml:"reaction_roles"`
	}

	// ReactionRole is the reaction role config
	ReactionRole struct {
		EmojiID   string `yaml:"emoji"`
		MessageID string `yaml:"message"`
		RoleID    string `yaml:"role"`
	}

	// Hub is a single faceit hub
	Hub struct {
		Format string `yaml:"format"`
		HubID  string `yaml:"hub_id"`
		RoleID string `yaml:"role_id"`
		//EmojiID   string `yaml:"emoji_id"`
		//MessageID string `yaml:"message_id"`
		//SkillLvl  int    `yaml:"skillvl"`
	}
)

func NewConfig() (*Config, error) {
	return nil, nil
}
