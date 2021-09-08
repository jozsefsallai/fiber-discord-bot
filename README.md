# Fiber Discord bot

This is a basic Discord bot that we mainly use on the Fiber Discord server to
manage administrative tasks (such as member approvals).

## Getting Started

**1. Clone the repo:**

```sh
git clone git@github.com:jozsefsallai/fiber-discord-bot.git
cd fiber-discord-bot
```

**2. Create a local config:**

```sh
cp .env.example .env
nvim .env
```

**3. Start the bot:**
```
go build
./fiber-discord-bot
```

or

```
go run main.go
```

## License

MIT.
