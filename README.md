# Drink Water Bitch

A Go application that sends hydration reminders to Google Chat because someone keeps forgetting to drink water.

Built to remind a coworker of mine to stay hydrated. if I have to hear one more dry ass organ story, I’m pushing their desk out of the window.

Inspired by [@drinkwaterslut](https://x.com/drinkwaterslut) on X (formerly Twitter). Some of the phrases in this list were taken from their X account because we missed them. 🙂


## Setup

1. Create a Google Chat webhook:
   - Open your Google Chat space
   - Click space name → **Manage webhooks**
   - Create a new webhook and copy the URL

2. Get the user ID:
   - Go read https://stackoverflow.com/a/58923385/9392409

3. Set environment variables:
   ```bash
   export GOOGLE_CHAT_WEBHOOK="your-webhook-url"
   export USER_ID="user-id-number"
   ```

4. Run:
   ```bash
   go run main.go
   ```

## Cron Setup

Ask ChatGPT

## How It Works

1. Picks a random phrase from `phrases.txt`
2. Sends it to Google Chat mentioning the user
3. Keeps your coworker's kidneys happy

## Disclaimer

All phrases are meant with love. Stay hydrated, bitches.
