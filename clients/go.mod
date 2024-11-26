module clients

go 1.23

require (
	common v0.0.0-00010101000000-000000000000
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
)

require github.com/technoweenie/multipartstreamer v1.0.1 // indirect

replace common => ../common
