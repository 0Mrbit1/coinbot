package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"sync"
	"time"
	"log"
)

type MessageDeletion struct {
	ChatId    int64
	MessageId int64
	ChatType string
}

var (
	deletionChannel = make(chan MessageDeletion, 100) // Buffered channel for pending deletions
	wg              sync.WaitGroup
)

func handleMessageDeletions(b *gotgbot.Bot) {

	for msg := range deletionChannel {
		wg.Add(1)
		go func(m MessageDeletion) {
			defer wg.Done()
			time.Sleep(60 * time.Second)
			if m.ChatType != "private"{
			_, err := b.DeleteMessage(m.ChatId, m.MessageId, nil)
			if err != nil {
				log.Printf("Failed to delete message: %s", err)
			} else {
				log.Printf("Message %d deleted successfully.", m.MessageId)
			}
		}
		}(msg)
	}
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	msg, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Welcome %s please chose from the list of commands to get started.", ctx.EffectiveChat.FirstName), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId) , ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}

	// Add the bot's message to the deletion channel
	return nil
}

func about_feyorra(b *gotgbot.Bot, ctx *ext.Context) error {
	reply_text := "The FaucetPay team is thrilled to announce creating a utility token of our own. This is a very exciting moment for us since we worked actively for a year to develop and launch the token!We would like to present Feyorra (FEY).FEY will be the utility token for all projects associated with Basilisk Entertainment S.R.L. On top of that, it comes with several important DeFi features. Check more about feyorra  here :"
	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Feyorra", Url: "https://feyorra.com"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}

func stake(b *gotgbot.Bot, ctx *ext.Context) error {
	reply_text := "There are several ways to stake Feyorra :"
	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Decentralized staking (INSTALL METAMASK)", Url: "https://dapp.feyorra.com"}},{
				{Text: "Pooled Staking", Url: "https://faucetpay.io/fey/pooled-staking"}},{
				{Text: "Help", Url: "https://feyorra.medium.com"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}


	// Add the bot's message to the deletion channel
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}

func stats(b *gotgbot.Bot, ctx *ext.Context) error {

	stats , err_api := fetchStatistics()
	
	if err_api != nil {
        fmt.Println("Error:", err_api)
        return nil
    }
	 // Print the struct fields to verify

	reply_text := stats

	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Refresh", CallbackData: "data_refresh"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	// Add the bot's message to the deletion channel
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}

func utility(b *gotgbot.Bot, ctx *ext.Context) error {

	reply_text := "Verify your stakes at Pasino.com and earn a share of the house edge! Pasino pays out 0.1% of the total betting volume to all verified stakers. Payouts are made in over 10 different cryptocurrencies. House edge-sharing is your gateway to an additional passive income."
	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Pasino", Url: "https://pasino.com/house-edge-sharing-feyorra"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	// Add the bot's message to the deletion channel
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}

func exchanges(b *gotgbot.Bot, ctx *ext.Context) error {

	reply_text := "You can trade Feyorra here:"
	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Uniswap", Url: "https://v2.info.uniswap.org/pair/0xb6e544c3e420154c2c663f14edad92737d7fbde5"}},{
				{Text: "FAUCETPAY-SWAP", Url: "https://faucetpay.io"}},{
				{Text: "PASINO-SWAP", Url: "https://pasino.com"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	// Add the bot's message to the deletion channel
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}

func contract(b *gotgbot.Bot, ctx *ext.Context) error {

	reply_text := "0xe8e06a5613dc86d459bc8fb989e173bb8b256072"
	msg, err := ctx.EffectiveMessage.Reply(b, reply_text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "EtherScan", Url: "https://etherscan.io/token/0xe8e06a5613dc86d459bc8fb989e173bb8b256072"},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	// Add the bot's message to the deletion channel
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(ctx.Message.MessageId), ChatType: ctx.EffectiveChat.Type}
	deletionChannel <- MessageDeletion{ChatId: ctx.EffectiveChat.Id , MessageId: int64(msg.MessageId), ChatType: ctx.EffectiveChat.Type}
	return nil
}