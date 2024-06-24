package main


import (
	"log"
	"time"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

// This bot demonstrates some example interactions with telegram callback queries.
// It has a basic start command which contains a button. When pressed, this button is edited via a callback query.
func main() {
	// Get token from the environment variable
	token := ""
	if token == "" {
		panic("TOKEN variable is empty")
	}
	// Create bot from environment value.
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("about_feyorra", about_feyorra))
	dispatcher.AddHandler(handlers.NewCommand("stats", stats))
	dispatcher.AddHandler(handlers.NewCommand("stake", stake))
	dispatcher.AddHandler(handlers.NewCommand("contract", contract))
	dispatcher.AddHandler(handlers.NewCommand("exchanges", exchanges))
	dispatcher.AddHandler(handlers.NewCommand("utility", utility))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("data_refresh"), startCB))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

