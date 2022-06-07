package bot

import (
	"SimpleModerationBot/config"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	currentConfig *config.ConfigStruct
)

type botStruct struct {
	botSession *discordgo.Session
}

func NewBot(config *config.ConfigStruct) (*botStruct, error) {
	currentConfig = config
	botSession, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	botSession.AddHandler(messageHandler)

	return &botStruct{botSession: botSession}, nil
}

func (bot *botStruct) Start() error {
	err := bot.botSession.Open()
	if err != nil {
		return err
	}

	log.Print("Bot is running!")

	return nil
}

func (bot *botStruct) Close() {
	log.Print("Closing bot...")
	bot.botSession.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}

	if blacklisted, msg := containsBlacklisted(m); blacklisted {
		log.Print("found blacklisted word \"" + *msg + "\"")
		s.ChannelMessageSend(m.ChannelID, "Stop using `"+*msg+"`. Are you 12?")
	}
}

func containsBlacklisted(m *discordgo.MessageCreate) (bool, *string) {
	msgLower := strings.ToLower(m.Content)

	for _, word := range currentConfig.BlackList {
		word = strings.ToLower(word)
		if match, _ := regexp.MatchString("\\b"+word+"\\b", msgLower); match {
			return true, &word
		}
	}

	return false, nil
}
