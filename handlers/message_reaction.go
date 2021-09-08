package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jozsefsallai/fiber-discord-bot/config"
)

func HandleMessageReaction(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	if reaction.MessageID != config.Config.Server.Approvals.RulesMessageID {
		return
	}

	if reaction.Emoji.Name != "✅" {
		return
	}

	guild := reaction.GuildID

	role, err := session.State.Role(guild, config.Config.Server.Approvals.ApprovedRoleID)
	if err != nil {
		log.Println(err)
		return
	}

	member, err := session.GuildMember(guild, reaction.UserID)
	if err != nil {
		log.Println(err)
		return
	}

	err = session.GuildMemberRoleAdd(guild, member.User.ID, role.ID)
	if err != nil {
		log.Println(err)
		return
	}
}
