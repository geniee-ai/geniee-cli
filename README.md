# geniee-cli

Geniee is a CLI based tool which allows you to use ChatGPT directly from terminal.

## Setup

1. Signup to https://geniee.io and generate a access token.

2. Install Geniee CLI

```shell
brew tap --cask geniee
brew install pgaijin66/geniee
```

3. Login using `geniee login`

4. Ask questions


## Available commands


```shell
$ geniee help

NAME:
   Geniee - ask any questions directly from terminal

USAGE:
   Geniee [global options] command [command options] [arguments...]

VERSION:
   v0.1.2

COMMANDS:
   ask, a      Ask questions
   login, l    Obtain and save credentials from cheesy web.
   version, v  Show version
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

WEBSITE: https://geniee.io

```




## Improvements

- [x] Add ability for user to manage their own token

- [x] Support for longer questions

- [x] Migrated to ChatGPT 3.5

- [ ] Make usage more interactive with follow up questions

- [ ] Implement caching for frequent similarly asked questions

- [ ] TBD



