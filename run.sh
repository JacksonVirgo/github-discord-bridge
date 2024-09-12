#!/bin/bash
docker build -t github-discord-bridge .
docker run -p 8080:8080 github-discord-bridge
