const fs = require('fs');
const path = require('path');

const env = process.env.NODE_ENV || 'development';
let config = {};

const envConfigPath = path.resolve(__dirname, env);

if (!fs.existsSync(envConfigPath) && env !== 'production') {
  throw new Error(`No config for NODE_ENV=${env}.`);
}

if (env !== 'production') {
  const envKeys = fs.readdirSync(envConfigPath);
  envKeys.map(f => path.basename(f, '.json')).forEach(key => {
    config[key] = require(`./${env}/${key}.json`);
  });
} else {
  const secretsConfigFile = '/var/secrets/fiber-bot.json';

  if (!fs.existsSync(secretsConfigFile)) {
    throw new Error('Coud not find secrets config file. The bot cannot be started.');
  }

  config = require(secretsConfigFile);
}

module.exports = config;
