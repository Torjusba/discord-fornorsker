package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/torjusba/discord-fornorsker/pkg/logging"
	"github.com/torjusba/discord-fornorsker/pkg/util"
	"github.com/torjusba/discord-fornorsker/pkg/wordlist"
)

func main() {
	logging.Log("Initializing word list...")

	wordlist.Init()

	logging.Log("Word list initialized")
	logging.Log("Initializing bot...")
	var Token string
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	discordSession, err := discordgo.New("Bot " + Token)
	if err != nil {
		logging.Error("Error creating Discord session", err)
		return
	}

	// Handle the MessageCreate event
	discordSession.AddHandler(handleReceivedMessage)
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	err = discordSession.Open()
	if err != nil {
		logging.Error("Error opening Discord session", err)
		return
	}

	logging.Log("The bot is now running. Press CTRL-C to exit.")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChan
	logging.Log("Exited cleanly")

}

func handleReceivedMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {

	// Don't answer ourselves
	if msg.Author.ID == session.State.User.ID {
		return
	}

	// Handle commands
	if strings.Contains(strings.ToLower(msg.Content), "!språkrådet") {
		command := strings.Split(strings.ToLower(msg.Content), " ")[1]
		switch command {
		case "hjelp":
			session.ChannelMessageSendReply(msg.ChannelID, util.HelpMessage, msg.MessageReference)
			return
		case "kildekode":
			session.ChannelMessageSendReply(msg.ChannelID, util.SourceCodeMessage, msg.MessageReference)
			return
		default:
			session.ChannelMessageSendReply(msg.ChannelID, util.UndefinedCommandMessage, msg.MessageReference)

		}
	}

	neededReplacements := wordlist.GetNeededReplacements(msg.Content)
	for word, replacement := range neededReplacements {
		reply := fmt.Sprintf("Heisann! %s \nDu skrev '%s', som er lånt fra engelsk. Jeg anbefaler deg å bruke det norske alternativet:\n%s", msg.Author.Mention(), word, replacement)
		session.ChannelMessageSendReply(msg.ChannelID, reply, msg.MessageReference)
	}
}
