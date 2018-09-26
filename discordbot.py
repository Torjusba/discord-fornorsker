import discord
from discord.ext.commands import Bot
from discord.ext import commands
import asyncio
import time

Client = discord.Client()
client = commands.Bot(command_prefix="#")

privatetoken = ""

with open("discordbot.token", "r") as f:
    privatetoken = f.read()

print(privatetoken)

#Show that bot is connected
@client.event
async def on_ready():
    print("Bot is online")
    print(client.user)

#Handle messages
@client.event
async def on_message(message):
    #Do not answer self
    if message.author == client.user:
        return()

    
    
    _message = message.content.lower()
    if _message == "rc":
        await client.send_message(message.channel, "lc")
    elif "gud bevare konge og fedreland" in _message:
        await client.send_message(message.channel, ":flag_bv: JA! :flag_bv:")
    elif "radio" in _message:
        await client.send_message(message.channel, "STERK OG KLAR!")
    elif _message == "cito":
        time.sleep(1)
        await client.send_message(message.channel, "this is how we do it down in Puerto Rico")
    elif "kontakt" in _message:
        _msg = _message.replace("kontakt", "forbindelse")
        await client.send_message(message.channel, 'Mente du: "' + _msg + '"?')


client.run(privatetoken)
