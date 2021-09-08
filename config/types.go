package config

type ApprovalsConfig struct {
	RulesChannelID string
	RulesMessageID string
	ApprovedRoleID string
}

type ServerConfig struct {
	ID        string
	Approvals *ApprovalsConfig
}

type BotConfig struct {
	Token string
}

type AppConfig struct {
	Bot    *BotConfig
	Server *ServerConfig
}
