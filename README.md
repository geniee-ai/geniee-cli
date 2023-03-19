![geniee](assets/geniee.png)


# Geniee

Geniee is a CLI-based tool that allows you to use ChatGPT directly from the terminal.

# Installation

## OSX 

```shell
brew tap pgaijin66/geniee
brew install --cask geniee
```

## Linux

```shell
wget https://github.com/geniee-ai/geniee-cli/releases/download/v0.1.2/geniee_0.1.2_linux_amd64
mv geniee_0.1.2_linux_amd64 /usr/local/bin/geniee
chmod +x /usr/local/bin/geniee
```

## Setup

1. Sign up for https://geniee.io and generate an access token. Log in to the web application, and navigate to Dashboard > Create Token (Note: Tokens are only shown once, so be sure to copy and save it somewhere safe.).

2. Install Geniee CLI according to the specifications of your operating system.

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


`geniee login`: Authenticate against Geniee web services and log in. This is a one-time step unless you have rotated your access token from the web. (Note: Here, you would provide the access token generated from Geniee web.)

`geniee ask` : Ask questions to ChatGPT using Geniee CLI. Before using this command, you will need to log in.

`geniee version` : Show current app version 

## Improvements

- [x] Add ability for user to manage their own token

- [x] Support for longer questions

- [x] Migrate to ChatGPT 3.5

- [ ] Make usage more interactive with follow up questions

- [ ] Implement caching for frequent similarly asked questions

- [ ] Ability to pass file as an argument and get response based on that.

- [ ] TBD


## Website

https://geniee.io


