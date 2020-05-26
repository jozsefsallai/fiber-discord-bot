# Fiber Discord bot

This is a basic Discord bot that we mainly use on the Fiber Discord server to manage administrative tasks (such as member approvals). It is build upon [discord-akairo](https://discord-akairo.github.io).

## Getting Started

**1. Clone the repo:**

```sh
git clone git@github.com:jozsefsallai/fiber-discord-bot.git
cd fiber-discord-bot
```

**2. Install the dependencies:**

```sh
npm i
```

**3. Create a local config:**

```sh
cp -r config/example config/development
cd config/development
vi bot.json
vi server.json
```

**4. Run in watch mode using nodemon:**

```sh
npm i -g nodemon
nodemon
```

## Production

In production environments, the configuration is loaded from `/var/secrets/fiber-bot.json`. To keep the bot running, use something like [pm2](https://npmjs.com/package/pm2).

## License

MIT.
