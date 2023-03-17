![geniee](assets/geniee.png)


# Geniee

Geniee is a CLI based tool which allows you to use ChatGPT directly from terminal.

## Setup

1. Signup to https://geniee.io and generate a access token. Log into web, and go to Dashboard > Create token (*Note: Token are only shown once so place copy and save it somewhere.*)

2. Install Geniee CLI

```shell
brew tap --cask geniee
brew install pgaijin66/geniee
```

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

- [ ] TBD


## Website

https://geniee.io


