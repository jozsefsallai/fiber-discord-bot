package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var Config *AppConfig

func loadDotenv() error {
	customPath := os.Getenv("CONFIG_PATH")
	if len(customPath) > 0 {
		err := godotenv.Load(customPath)

		if err == nil {
			return nil
		}
	}

	exepath, err := os.Executable()
	if err != nil {
		return err
	}

	exedir := filepath.Dir(exepath)
	err = godotenv.Load(filepath.Join(exedir, ".env"))
	return err
}

func Load() {
	loadDotenv()

	if os.Getenv("BOT_TOKEN") == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	Config = &AppConfig{
		Bot: &BotConfig{
			Token: os.Getenv("BOT_TOKEN"),
		},

		Server: &ServerConfig{
			ID: os.Getenv("SERVER_ID"),

			Approvals: &ApprovalsConfig{
				RulesChannelID: os.Getenv("RULES_CHANNEL_ID"),
				RulesMessageID: os.Getenv("RULES_MESSAGE_ID"),
				ApprovedRoleID: os.Getenv("APPROVED_ROLE_ID"),
			},
		},
	}
}
