# vhdbot - Valheim discord bot

## Overview
This is a quickly put together Discord bot to control your
dedicated Valheim server without the need to ssh into the linux machine

In other words, you can give access to people who you don't want ssh-ing into your server

It is meant to be used with linux-only (well at least at this point :))

## Pre-requisites
* Discord app + bot configured -> https://discord.com/developers/applications
* Discord bot token - look at 1st pre-req
* Everything else Discord :)

## Functionality

Right now it's meant to run on the same server where the dedicated Valheim server is
And in fact, since I was lazy, I hardcoded the path to the server start/stop script too :)

Once you build the bot, the following flags can be passed to the binary
* -t - bot token (mandatory)

all other flags are optional
* -c allows you to configure your bot channel command character
* -l gives you a list of all your channels and their IDs
* -ac allows you to specify the channel ID for sending single messages to (only works with -m)
* -m flag contains data for ad-hoc messages to the channel from -ac (only works with -ac)

I.e.
vhdbot -t "ODa0MkU0MjM5NTk2Nzg5NzWY.YH_q8A.qClWB0h4IGovsB0J1HcuCIhWPKc" -ac 123426540469387274 -m "Server went down!!!"

In the channel where your bot listens, you can do a few basic commands (default character is !)
!start - start your Valheim dedicated server
!stop - stop it
rest is not yet implemented since I'm lazy :)

## NB

DOES NOT HAVE AUTHORIZATION!!!
Every user on the channel where the bot listens will be able to issue the commands to it
Recommend you either run it on a server with friends or in a private/restricted channel
