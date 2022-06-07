package main

import (
	"SimpleModerationBot/bot"
	"SimpleModerationBot/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := config.LoadConfig()

	botSession, err := bot.NewBot(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = botSession.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

	waitForInterrupt()

	botSession.Close()
}

func waitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
