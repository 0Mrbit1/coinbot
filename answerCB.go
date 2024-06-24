package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery

	stats , err_api := fetchStatistics()
	
	if err_api != nil {
        fmt.Println("Error:", err_api)
        return nil
    }
	 // Print the struct fields to verify

	reply_text := stats

	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "Success",
	})
	if err != nil {
		return fmt.Errorf("failed to answer start callback query: %w", err)
	}

	_, _, err = cb.Message.EditText(b, reply_text, &gotgbot.EditMessageTextOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Refresh", CallbackData: "data_refresh"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to edit start message text: %w", err)
	}
	return nil
}