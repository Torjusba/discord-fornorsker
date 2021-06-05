package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/torjusba/discord-fornorsker/pkg/logging"
)

func main() {
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

	session.ChannelMessageSendReply(msg.ChannelID, "ok", msg.MessageReference)
}
