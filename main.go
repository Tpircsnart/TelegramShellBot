package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"telegram-shell-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard tgbotapi.ReplyKeyboardMarkup
var appSetting models.Setting

func main() {
	bot, err := tgbotapi.NewBotAPI(appSetting.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = appSetting.IsDebug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		isValidWord := false
		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
			isValidWord = true
		case "reset":
			err = SetButtons()
			if err != nil {
				log.Println(err)
				msg.Text = err.Error()
				break
			}
			msg.ReplyMarkup = numericKeyboard
			isValidWord = true
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			isValidWord = true
		}

		for _, v := range appSetting.Buttons {
			if v == msg.Text {
				isValidWord = true

				cmd := exec.Command("bash", appSetting.ShellLocation+v+".sh")
				output, err := cmd.CombinedOutput()
				if err != nil {
					log.Println(err)
					msg.Text = err.Error()
					break
				}
				msg.Text = string(output)
			}
		}

		if !isValidWord {
			msg.Text = "not valid word"
		}

		if len(msg.Text) > 4096 { // Telegram limit
			fmt.Println("yes")
			msg.Text = fmt.Sprintf("%s\n...\n...\n...\nOVER TELEGRAM TEXTSIZE", msg.Text[0:4000])
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func init() {
	err := error(nil)
	err = SetSettings()
	if err != nil {
		panic(err)
	}
	if appSetting.Token == "" {
		panic("token must be set")
	}
	if appSetting.RowButtonCount <= 0 {
		panic("line button count must be set over 0")
	}

	err = SetButtons()
	if err != nil {
		panic(err)
	}
}

func SetSettings() (err error) {
	jsonFile, err := os.Open("./env/settings.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	dataBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(dataBytes, &appSetting)
	return
}

func SetButtons() (err error) {
	files, err := os.ReadDir(appSetting.ShellLocation)
	if err != nil {
		return
	}

	shellFiles := []string{}
	for _, v := range files {
		filename := v.Name()
		if filepath.Ext(appSetting.ShellLocation+filename) == ".sh" {
			shellFiles = append(shellFiles, filename[:len(filename)-3])
		}
	}
	shellFilesLen := len(shellFiles)
	if shellFilesLen == 0 {
		err = errors.New("empty shell files")
		return
	}

	wholeButtons := [][]tgbotapi.KeyboardButton{}
	rowButtons := []tgbotapi.KeyboardButton{}

	for i, v := range shellFiles {
		appSetting.Buttons = append(appSetting.Buttons, v)
		rowButtons = append(rowButtons, tgbotapi.NewKeyboardButton(v))
		if (i+1)%appSetting.RowButtonCount == 0 || i+1 == shellFilesLen {
			wholeButtons = append(wholeButtons, rowButtons)
			rowButtons = []tgbotapi.KeyboardButton{}
		}
	}
	numericKeyboard = tgbotapi.NewReplyKeyboard(wholeButtons...)
	return
}
