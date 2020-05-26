const { AkairoClient } = require('discord-akairo');
const config = require('./config');

const {
  cacheRulesMessage,
  handleApprovalReaction,
  handleMemberJoin
} = require('./handlers/approval');

class FiberBotClient extends AkairoClient {
  constructor() {
    super({
      ownerID: config.bot.owners
    }, {
      disableEveryone: true
    });
  }
}

const bot = new FiberBotClient();

bot.on('ready', function () {
  cacheRulesMessage(bot);
});

bot.on('guildMemberAdd', member => {
  handleMemberJoin(member);
});

bot.on('messageReactionAdd', (reaction, user) => {
  handleApprovalReaction(reaction, user);
});

bot
  .login(config.bot.token)
  .then(function () {
    console.log('Bot started.');
  })
  .catch(console.error);
