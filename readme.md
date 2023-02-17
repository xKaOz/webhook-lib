This is a simple Go library for sending messages with embeds to a Discord webhook.

`To install this package, run the following command:`

- go get github.com/xKaOz/webhook-lib

`To import this package, copy this:`

- import "github.com/xKaOz/webhook-lib"

`Example of how to send a webhook:`

 
```go

package main

import (
	"log"

	"github.com/xKaOz/webhook-lib"
)

func main() {
	webhook := &webhook_lib.Webhook{
		URL:      "https://discord.com/api/webhooks/1234567890/abcdefghijklm",
		Username: "MyBot",
		Avatar:   "https://example.com/avatar.png",
		Content:  "Hello, world!",
		Embeds: []*webhook_lib.Embed{
			{
				Title: "Embed Title",
				Description: "Embed Description",
				Color: 0xFF0000,
				Fields: []*webhook_lib.EmbedField{
					{
						Name:  "Field 1",
						Value: "Value 1",
					},
					{
						Name:  "Field 2",
						Value: "Value 2",
					},
				},
			},
		},
	}

	if err := webhook.Send(); err != nil {
		log.Fatalf("Failed to send webhook: %v", err)
	}
}
```
