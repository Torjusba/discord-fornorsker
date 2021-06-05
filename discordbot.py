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

bigserver_whitelist = list([
    "abort",
	"access",
	"allround",
	"aircondition",
	"audition",
	"backlog",
	"backslash",
	"backup",
	"barbecue",
	"benchmark",
	"boom",
	"booster",
	"boote",
	"browser",
	"bug",
	"build",
	"button",
	"cap",
	"case",
	"chat",
	"chatte",
	"chip",
	"chips",
	"clickbait",
	"clutch",
	"connector",
	"controller",
	"cookie",
	"crew",
	"debug",
	"design",
	"driver",
	"equalizer",
	"fan",
	"franchise",
	"gate",
	"green",
	"guide",
	"hardware",
	"highlight",
	"host",
	"hoste",
	"hosting",
	"laptop",
	"link",
	"live",
	"makeup",
	"mashup",
	"match",
	"matche",
	"multitasking",
	"mute",
	"netbook",
	"offline",
	"online",
	"paring",
	"patch",
	"pitch",
	"polish",
	"poster",
	"printe",
	"rack",
	"release",
	"roaming",
	"sample",
	"sampling",
	"scanne",
	"scanner",
	"server",
	"skins",
	"slash",
	"sound",
	"standby",
	"support",
	"tab",
	"talkshow",
	"tape",
	"themes",
	"thumbnail",
	"tights",
	"time",
	"touchpad",
	"trigge",
	"trigger",
	"tutorial",
	"twitter",
	"webserver",
	"whiteboard",
	"widescreen",
	"wizard",
    "@"
])

smallserver_whitelist = list([
    "abort",
    "gate",
    "host",
    "hoste",
    "hosting",
    "live",
    "tape",
    "time",
    "twitter",
    "@"
])

def check_hit(_hitrate):
	roll = random.randint(0,100)
	return(roll<=_hitrate)

def check_respond(msg, content):
    global bigserver_whitelist
    global smallserver_whitelist

    if "språkrådet" in content:
        return True

    # Tilpass til store og små servere
    if msg.server.large:
        rate = 20
        whitelist = bigserver_whitelist
    else:
        whitelist = smallserver_whitelist
        rate = 100
    
    #Drit i irriterende ord
    for w in whitelist:
        if w in content:
            return False

    # Treffhyppighet
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
    global smallserver_whitelist
    global bigserver_whitelist
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
        if message.server.large:
            whitelist = bigserver_whitelist
        else:
            whitelist = smallserver_whitelist
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
