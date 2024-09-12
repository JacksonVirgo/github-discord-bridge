# Github Discord Bridge

## Installation

### Prerequisites

- Ensure Go is installed, you can download Go [here](https://golang.org/doc/install) (v1.23.1 or later)
- Ensure Docker is installed on your machine. You can download Docker [here](https://www.docker.com/get-started).

## Steps

1. Clone the repository `git clone https://github.com/JacksonVirgo/github-discord-bridge`
2. Navigate to the project directory `cd github-discord-bridge`
3. Build the docker image `docker build -t github-discord-bridge .`
4. Run the docker image `docker run -p 8080:8080 github-discord-bridge`
