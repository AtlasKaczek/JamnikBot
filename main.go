package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"stankryj/JamnikBot/aplikacja"
	"strconv"
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
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!jamnik" {
		jamnik, err := aplikacja.GetRandomJamnik("https://www.reddit.com/r/Dachshund.json")
		if err != nil {
			fmt.Println(err)
		}
		_, merr := s.ChannelMessageSend("915909449829482498", strconv.Itoa(jamnik.GetChildrenLen()))
		if merr != nil {
			fmt.Println(merr)
		}
	}
	if m.Content == "!halo" {
		_, err := s.ChannelMessageSend("915909449829482498", "halo!")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("MessegeSend: halo!\n")
		}
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
