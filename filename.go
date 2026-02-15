package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
)

var (
	[span_0](start_span)appID   int32  = 25742938                           // Tuhada API_ID[span_0](end_span)
	[span_1](start_span)appHash string = "b35b715fe8dc0a58e8048988286fc5b6" // Tuhada API_HASH[span_1](end_span)
	[span_2](start_span)token   string = "7623679464:AAGqdslPgtOzrAtycf6iuuDGPAJZCw4vJR0" // Tuhada Bot Token[span_2](end_span)
)

func main() {
	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
	})
	if err != nil {
		log.Fatal(err)
	}

	client.LoginBot(token)

	fmt.Println(">> Bot is running... Send /generate to your bot.")

	client.OnMessage(func(ctx *telegram.NewMessage) error {
		if ctx.Text() == "/start" || ctx.Text() == "/generate" {
			ctx.Reply("Gogram String Session generate karan layi apna Phone Number bhejo (with country code, e.g., +919876543210):")
			return nil
		}

		// Phone number handle karna
		if strings.HasPrefix(ctx.Text(), "+") {
			phone := ctx.Text()
			
			// Nawa user client create karna login layi
			userClient, _ := telegram.NewClient(telegram.ClientConfig{
				AppID:   appID,
				AppHash: appHash,
			})

			sentCode, err := userClient.AuthSendCode(phone)
			if err != nil {
				ctx.Reply("Error: " + err.Error())
				return nil
			}

			ctx.Reply("OTP mil gaya hai? Kripya OTP is format vich bhejo: `code:12345` (code de baad apna OTP likho)")
			
			// Simple handler for OTP (Note: Real bot vich state management chahidi hundi hai)
			client.OnMessage(func(otpCtx *telegram.NewMessage) error {
				if strings.HasPrefix(otpCtx.Text(), "code:") {
					code := strings.TrimPrefix(otpCtx.Text(), "code:")
					_, err := userClient.AuthSignIn(phone, sentCode.PhoneCodeHash, code)
					if err != nil {
						otpCtx.Reply("Invalid OTP: " + err.Error())
						return nil
					}

					session, _ := userClient.ExportSession()
					otpCtx.Reply("✅ Tuhadi Gogram String Session:\n\n`" + session + "`\n\nIsnu copy karke safe rakho.")
				}
				return nil
			})
		}
		return nil
	})

	client.Idle()
}
