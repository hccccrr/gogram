package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var (
	appID   int32  = 25742938
	appHash string = "b35b715fe8dc0a58e8048988286fc5b6"
	token   string = "7623679464:AAGqdslPgtOzrAtycf6iuuDGPAJZCw4vJR0"
)

func main() {
	// BOT CLIENT
	bot, err := tg.NewClient(tg.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: tg.LogInfo,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := bot.LoginBot(token); err != nil {
		log.Fatal("Bot login failed:", err)
	}

	fmt.Println("🤖 Bot online. Use terminal for session generation.")

	// ===== USER SESSION PART (TERMINAL BASED – SAFE & CLEAN) =====

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📞 Enter phone number (+91...): ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	user, err := tg.NewClient(tg.ClientConfig{
		AppID:   appID,
		AppHash: appHash,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := user.Connect(); err != nil {
		log.Fatal(err)
	}

	sentCode, err := user.AuthSendCode(
		phone,
		appID,
		appHash,
		&tg.CodeSettings{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("🔑 Enter OTP: ")
	otp, _ := reader.ReadString('\n')
	otp = strings.TrimSpace(otp)

	_, err = user.AuthSignIn(
		phone,
		sentCode.PhoneCodeHash,
		otp,
		tg.EmailVerification{},
	)
	if err != nil {
		log.Fatal("Login failed:", err)
	}

	session := user.ExportSession()

	fmt.Println("\n✅ GOGRAM STRING SESSION:\n")
	fmt.Println(session)

	bot.Idle()
}
