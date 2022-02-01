package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func goDotEnvVariable(key string) string {

	// load .env file
	errLoad := godotenv.Load(".env")

	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
	ChatID   int64  `json:"chat_id"`
}

type Info struct {
	Info InfoVal `json:"info"`
}

type InfoVal struct {
	Link string `json:"one_time_login_link"`
}

func main() {
	bot, err := tgbotapi.NewBotAPI(goDotEnvVariable("TOKEN_TLG"))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.

			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "great job")
			link := return_link(msg.ChatID) // If the message was open, add a copy of our numeric keyboard.
			if time.Now().UTC().Hour() <= 19 && time.Now().UTC().Hour() >= 7 {
				fmt.Println(time.Now().UTC().Hour())
				switch strings.ToUpper(update.Message.Text) {
				case "LOGIN":
					msg.Text = link
				case "/START":
					msg.Text = "enter your password"
				default:
					msg.Text = "nope"
				}

				// Send the message.
				if _, err = bot.Send(msg); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
