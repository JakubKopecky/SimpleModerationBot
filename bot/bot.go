package bot

import (
	"SimpleModerationBot/config"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
)

func Start() {

	goBot, err := discordgo.New("Bot " + config.LoadedConfiguration.Token)
	if err != nil {
		log.Panic(err)
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		log.Panic(err)
		return
	}

	BotId = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		log.Panic(err)
		return
	}
	log.Print("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if blacklisted, msg := containsBlacklisted(m); blacklisted {
		log.Print("found blacklisted word \"" + *msg + "\"")
		_, _ = s.ChannelMessageSend(m.ChannelID, "Stop using `"+*msg+"`. Are you 12?")
	}
}

func containsBlacklisted(m *discordgo.MessageCreate) (bool, *string) {
	msgLower := strings.ToLower(m.Content)

	for _, word := range config.LoadedConfiguration.BlackList {
		word = strings.ToLower(word)
		if match, _ := regexp.MatchString("\\b"+word+"\\b", msgLower); match {
			return true, &word
		}
	}

	return false, nil
}
