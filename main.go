package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"stankryj/JamnikBot/aplikacja"

	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!jamnik" {
		jamn, err := aplikacja.GetJamnikObj()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(jamn.GetImagesIndexList()))

		_, merr := s.ChannelMessageSend(m.ChannelID, jamn.GetImageURL(jamn.GetImagesIndexList()[rand.Intn(len(jamn.GetImagesIndexList()))]))
		if merr != nil {
			fmt.Println(merr)
		} else {
			fmt.Printf("MessegeSend: Image send!\n")
		}
	}
	if m.Content == "!halo" {
		_, err := s.ChannelMessageSend(m.ChannelID, "halo!")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("MessegeSend: halo!\n")
		}
	}
	// Add new command and reddit to the list - TODO
	if m.Content[0:4] == "!add" {
		res, url, err := aplikacja.GetCMDvariables(m.Content)
		if err != nil {
			fmt.Println(err)
		}
		_, merr := s.ChannelMessageSend(m.ChannelID, res+" "+url)
		if merr != nil {
			fmt.Println(err)
		}
		fmt.Printf("MessegeSend: ADD\n")
	}
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
