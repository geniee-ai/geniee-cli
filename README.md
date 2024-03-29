![geniee](assets/geniee.png)


# Geniee

Geniee is a CLI-based tool that allows you to use ChatGPT directly from the terminal.

# Installation

## OSX 

```shell
brew tap pgaijin66/geniee
brew install --cask geniee
```

For OSX users:

When you first run `geniee` command, apple security will pick it up and prompt you. To continue using it, you will have to allow it in security and privacy settings.

Steps to allow:

1. Go to System Settings > Privacy and Security2. 

2. Scroll below until you find the prompt. Click allow

3. Run `geniee` command again, and click 'open'. This is a one time thing and you won't have to do it again.

## Linux

```shell
wget https://github.com/geniee-ai/geniee-cli/releases/download/v0.1.2/geniee_0.1.2_linux_amd64
mv geniee_0.1.2_linux_amd64 /usr/local/bin/geniee
chmod +x /usr/local/bin/geniee
```

## Linux

```shell
wget https://github.com/geniee-ai/geniee-cli/releases/download/v0.1.2/geniee_0.1.2_linux_amd64
mv geniee_0.1.2_linux_amd64 /usr/local/bin/geniee
chmod +x /usr/local/bin/geniee
```

## Windows

We are working on making geniee available for windows user via Chocolately in the mean time you can take advantage of geniee using WSL feature.

1. Install WSL ( ubuntu ) on windows. Please follow this guide to install WSL ( https://pureinfotech.com/install-windows-subsystem-linux-2-windows-10 )

2. Once installed, you would need to reboot your PC

3. Once rebooted, search for WSL and open it.

4. It will setup ubuntu ( default distro for WSL ) and ask you to setup username / password

5. Once completed, follow the installation procedure for Linux shown above to install and run geniee.

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


