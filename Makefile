default: discordbot
discordbot:
	go build -o bin/discordbot cmd/discordbot/main.go

wordlist-generator:
	go build -o bin/wordlist-generator cmd/wordlist-generator/main.go