import discord
from discord.ext.commands import Bot
from discord.ext import commands
import asyncio
import time
import random

from privatetoken import getToken

from pdparser import PDParser

def recognize(_word, _importlist, _avloeyserlist):
    _alt = ""
    if _word in _importlist:
        _alt = _avloeyserlist[_importlist.index(_word)]
    return _alt

def get_alternative(_word, _importlist, _avloeyserlist):
    return _avloeyserlist[_importlist.index(_word)]
    



pdparser = PDParser('https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/')
Client = discord.Client()
client = commands.Bot(command_prefix="#")

privatetoken = getToken()

whitelist = list([
    "abort",
    "gate",
    "host",
    "hoste",
    "hosting",
    "live",
    "tape",
    "time",
    "twitter"
])

def check_hit(_hitrate):
	roll = random.randint(0,100)
	return(roll<=_hitrate)

def check_respond(msg, content):
    global whitelist
    if "språkrådet" in content:
        return True
    
    #Drit i irriterende ord
    for w in whitelist:
        if w in content:
            return False

    rate = 100
    if msg.server.large:
        rate = 20
    if check_hit(rate):
        return True
    return False


#Vis at programvareagenten er tilkoblet
@client.event
async def on_ready():
    print("Bot is online")
    print(client.user)

#Handle messages
@client.event
async def on_message(message):
    global whitelist
    #Do not answer self
    if message.author == client.user:
        return()

    _message = message.content.lower()

    if _message == "#språkrådet hjelp":
        await client.send_message(message.channel, "Svarer på 20% av importord på store servere, 100% på små.")
        await client.send_message(message.channel, "https://www.sprakradet.no/sprakhjelp/Skriverad/Avloeysarord/")
    
    if _message == "#språkrådet kanalstørrelse":
        await client.send_message(message.channel, "Stor kanal?: {}".format(message.server.large))

    if _message == "#språkrådet unntak":
        await client.send_message(message.channel, "Unntak fra retting: {}".format(whitelist))

    if not check_respond(message, _message):
        return()

    _arr = _message.split(" ")    

#    for w in _arr:
#        if (w=="yeet"):
#            _message = "Du brukte ordet 'yeet'. Dette er et lånord, og bør erstattes med å bare dø"
#            await client.send_message(message.channel, _message)
    
    for w in pdparser.importlist:
        if w in _message:
            if w == "nan": continue
            alt = get_alternative(w, pdparser.importlist, pdparser.avloeyserlist)
            _msgtosend = "Du brukte ordet '{}'. Dette er et lånord, og bør erstattes med et av følgende gode norske alternativer: \n {}".format(w, alt)
    await client.send_message(message.channel, _msgtosend)
            

client.run(privatetoken)
