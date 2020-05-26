const config = require('../config');

async function handleMemberJoin(member) {
  if (member.guild.id !== config.server.id) {
    return;
  }

  try {
    const dm = await member.createDM();

    await dm.send({
      embed: {
        title: 'Welcome to Fiber!',
        description: `Welcome to Fiber! To get verified, react with the :white_check_mark: emoji in <#${config.server.approvals.rulesChannelID}>`
      }
    });
  } catch (err) {
    console.error(err);
  }
}

async function cacheRulesMessage(bot) {
  const rulesChannel = await bot.channels.fetch(config.server.approvals.rulesChannelID);
  if (!rulesChannel) {
    return;
  }

  try {
    await rulesChannel.messages.fetch(config.server.approvals.rulesMessageID);
  } catch (err) {
    console.error(err);
  }
}

async function handleApprovalReaction(reaction, user) {
  if (reaction.message.id !== config.server.approvals.rulesMessageID) {
    return;
  }

  if (reaction.emoji.name !== '✅') {
    return;
  }

  const { guild } = reaction.message.channel;

  try {
    const role = await guild.roles.fetch(config.server.approvals.approvedRoleID);
    if (!role) {
      return;
    }

    const member = await guild.members.fetch(user.id);
    await member.roles.add(role);
  } catch (err) {
    console.error(err);
  }
}

module.exports = {
  cacheRulesMessage,
  handleMemberJoin,
  handleApprovalReaction
};
