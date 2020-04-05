package logger

import (
	"net/http"
	"net/url"
)

// defaultAPIToken is a Telegram Bot API access token used for logging notification(the token below is a dummy one!).
const defaultAPIToken = "427627935:AAFYF7PzAWfQBFiAEqeijbpu5TOJsydEbSo"

// defaultTargetUser is a Telegram unique user ID that the logging notification will be sent to(the user ID below is a dummy one!).
const defaultTargetUser = "254169872"

type TelegramLogger struct {
	APIToken string
	TargetUser string
}

// Init initializes TelegramLogger object and assigns some default values to its fields.
func (t *TelegramLogger) Init() *TelegramLogger {
	t.APIToken = defaultAPIToken
	t.TargetUser = defaultTargetUser

	return t
}

// Log method calls Telegram API and sends the log message to TargetUser.
func (t *TelegramLogger) Log(data string) {
	r, _ := http.Get("https://api.telegram.org/bot" + t.APIToken + "/sendMessage?chat_id=" + t.TargetUser + "&text=" + url.QueryEscape(data))
	r.Body.Close()
}
