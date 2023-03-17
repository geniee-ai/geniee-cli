![geniee](assets/geniee.png)


# Geniee

Geniee is a CLI based tool which allows you to use ChatGPT directly from terminal.

# Installation

## OSX 

```shell
brew tap --cask geniee
brew install pgaijin66/geniee
```

## Linux

```shell
curl https://github.com/geniee-ai/geniee-cli/releases/download/v0.1.2/geniee_0.1.2_linux_amd64
mv geniee_0.1.2_linux_amd64 /usr/local/bin/geniee
chmod +x /usr/local/bin/geniee
```

## Setup

1. Signup to https://geniee.io and generate a access token. Log into web, and go to Dashboard > Create token (*Note: Token are only shown once, so copy and save it somewhere safe.*)

2. Install Geniee CLI



3. Login using `geniee login`

4. Ask questions using `geniee ask`

eg:
```shell
$ geniee ask "what is a Dockerfile?

You request has been processed.

Here is your response.

Note: It might take a while to parse response in the terminal. Please have patience.



A Dockerfile is a text file that contains commands that are used to build a Docker image. It is a script that describes the steps necessary to create a Docker container. The Dockerfile contains instructions for building a specific image, including the base image, any additional packages or libraries that are required, and any configuration options that need to be set. Essentially, it is the blueprint for creating a Docker container or image.

```


## Available commands


`geniee login`: Authenticated against geniee web services and login. This is one time step, unless you rotated your access token from web. (*Note: Here you would provide the access token generated from geniee web. *)

`geniee ask` : Ask questions to ChatGPT using geniee CLI. You would need to login before using this command.

`geniee version` : Show current app version 

## Improvements

- [x] Add ability for user to manage their own token

- [x] Support for longer questions

- [x] Migrated to ChatGPT 3.5

- [ ] Make usage more interactive with follow up questions

- [ ] Implement caching for frequent similarly asked questions

- [ ] Pass file as an input option and get response based on that.

- [ ] TBD


## Website

https://geniee.io


