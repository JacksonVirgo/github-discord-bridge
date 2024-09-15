# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2024-09-16

### Added

- Basic command `!sync-tags` exists to aid syncing before the full github integration
- When thread tags change, the appropriate github labels will also change
- When threads are created, any attached tags will be sent as issue labels
- Messages in linked threads create a github issue comment
- Created threads in the linked channel creates a github issue
- Basic docker build
- Basic Go project
- This repository

[Unreleased]: https://github.com/JacksonVirgo/go-github-discord-bridge
[0.1.0]: https://github.com/JacksonVirgo/github-discord-bridge/releases/tag/0.1.0
