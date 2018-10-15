import discord
from discord.ext.commands import Bot
from discord.ext import commands
import asyncio
import time
import random

from pdparser import PDParser

def recognize(_word, _importlist, _avloeyserlist):
    _alt = ""
    if _word in _importlist:
        _alt = _avloeyserlist[_importlist.index(_word)]
    return _alt

pdparser = PDParser('https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/')
Client = discord.Client()
client = commands.Bot(command_prefix="#")

privatetoken = "NDk0NTQ3NzcxNjI2MzU2NzU4.Dp5-ZQ.KWnfuA7fgEfX48eeWBlIqcxEhY4"

def check_hit(_hitrate):
	roll = random.randint(0,100)
	return(roll<=_hitrate)


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
    
    if _message == "@språkrådet hjelp":
        await client.send_message(message.channel, "Svarer på 50% av importord. https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/'")
    
    if not check_hit(50):
        return()

    _arr = _message.split(" ")    

    for w in _arr:
        alt = recognize(w, pdparser.importlist, pdparser.avloeyserlist)
        if len(alt)>0:
            _message = "Du brukte ordet '{}'. Dette er et lånord, og bør erstattes med et av følgende gode norske alternativer: \n {}".format(w, alt)
            await client.send_message(message.channel, _message)
client.run(privatetoken)
