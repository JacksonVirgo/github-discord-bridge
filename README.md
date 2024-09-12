# Github Discord Bridge

This discord bot serves as a bridge between Discord forum channels and a Github repository's issues. Enabling more effective issue management with teams using both platforms.

View the changelog [here](CHANGELOG.md)

## Features

- Creating discord forum threads automatically creates a Github Issue
  - Tags are currently not supported
- Messages sent in linked threads will be sent as comments on the appropriate Github issue

## Plans

- Github to Discord linking
- Thread tags
- User assignments
- Allowing images (such as attachments, or even user icons for clarity)
- Issue lock/unlocks
- Issue opens/closes
- Deleting issues

## Installation

### Prerequisites

- Ensure Go is installed, you can download Go [here](https://golang.org/doc/install) (v1.23.1 or later)
- Ensure Docker is installed on your machine. You can download Docker [here](https://www.docker.com/get-started).

### Steps

1. Clone the repository `git clone https://github.com/JacksonVirgo/github-discord-bridge`
2. Navigate to the project directory `cd github-discord-bridge`
3. Build the docker image `docker build -t github-discord-bridge .`
4. Run the docker image `docker run -p 8080:8080 github-discord-bridge`
5. Create a discord bot account [here](https://discord.com/developers/applications?new_application=true)
   - Required Intents: Presence, Message Content
   - Invite it to your server `https://discord.com/api/oauth2/authorize?client_id=APPLICATION_ID&permissions=0&scope=bot`

### Setup

1. Copy the `.env.example` into `.env` e.g. `cp .env.example .env`
2. Fill out the appropriate information, more information is listed below this section

## Environment Variables

- DISCORD_TOKEN -> Discord developer bot page "Settings->bot->reset token"
- DISCORD_CHANNEL_ID -> In the Discord server, create a forum channel and right-click (RMB) to copy the channel ID (developer settings must be turned on for this).
- GITHUB_TOKEN
  1. Preface: Make sure you're creating these on the account that will be posting on Github. Such as a new account used primarily (and obviously) as a bot account
  2. [New Fine-grained Personal Access Token](https://github.com/settings/personal-access-tokens/new) or follow these steps: Settings -> Developer settings -> Personal access tokens -> Fine-grained tokens -> Generate new token.
  3. In the "Repository access" section, select "Only select repositories" and choose the specific repositories you need access to.
  4. In the "Permissions" section, click on "Repository permissions" and set "Issues" to "Read & Write".
  5. Generate and copy the personal access token.
- GITHUB_USERNAME -> The user the repository is under
- GITHUB_REPO -> The repository name (in its URL)
