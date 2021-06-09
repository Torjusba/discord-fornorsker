package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/torjusba/discord-fornorsker/pkg/analysis"
	"github.com/torjusba/discord-fornorsker/pkg/logging"
	"github.com/torjusba/discord-fornorsker/pkg/responses"
	"github.com/torjusba/discord-fornorsker/pkg/wordlist"
)

func main() {
	var token string
	var mode string
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&mode, "w", "simple", "Wordlist mode")
	flag.Parse()

	logging.Log("Initializing word list...")
	wordlist.Init(wordlist.Mode(mode))
	logging.Log("Word list initialized")

	logging.Log("Initializing bot...")
	discordSession, err := discordgo.New("Bot " + token)
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

	logging.Log("Bot initialized")
	logging.Log("The bot is now running. Press CTRL-C to exit.")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalCh
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
			session.ChannelMessageSendReply(msg.ChannelID, responses.HelpMessage, msg.MessageReference)
			return
		case "kildekode":
			session.ChannelMessageSendReply(msg.ChannelID, responses.SourceCodeMessage, msg.MessageReference)
			return
		default:
			session.ChannelMessageSendReply(msg.ChannelID, responses.UndefinedCommandMessage, msg.MessageReference)
			return
		}
	}

	// Actual borrowed word replacement
	borrowedWordsFound := analysis.WhichWordsAreInSentence(msg.Content, wordlist.GetBorrowedWords())
	for _, word := range borrowedWordsFound {
		replacement := wordlist.GetReplacement(word)
		reply := fmt.Sprintf("Heisann! %s \nDu skrev '%s', som er lånt fra engelsk. Jeg anbefaler deg å bruke det norske alternativet:\n%s", msg.Author.Mention(), word, replacement)
		session.ChannelMessageSendReply(msg.ChannelID, reply, msg.MessageReference)
	}
}
