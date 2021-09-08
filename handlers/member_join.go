package handlers

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jozsefsallai/fiber-discord-bot/config"
)

func generateWelcomeEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Welcome to Fiber!",
		Description: fmt.Sprintf(
			"Welcome to Fiber! To get verified, react with the :white_check_mark: emoji in <#%s>.\n"+
				"If the automatic approval does not work, please contact a maintainer privately.",
			config.Config.Server.Approvals.RulesChannelID,
		),
	}
}

func HandleMemberJoin(session *discordgo.Session, member *discordgo.GuildMemberAdd) {
	if member.User.Bot {
		return
	}

	channel, err := session.UserChannelCreate(member.User.ID)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = session.ChannelMessageSendEmbed(channel.ID, generateWelcomeEmbed())
	if err != nil {
		log.Println(err)
		return
	}
}
