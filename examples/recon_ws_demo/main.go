package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/ericyeungcode/caliber"
)

func main() {
	wsURL := "wss://betaws.bitexch.dev"

	client := caliber.NewWSClient(wsURL, &tls.Config{
		InsecureSkipVerify: true, // ⚠️ only for dev / testing
	})

	client.OnOpen = func() {
		var subscription = map[string]any{
			"type":        "subscribe",
			"instruments": []string{"BTC-USDT-PERPETUAL"},
			"channels":    []string{"order_book.10.10"},
			"interval":    "raw",
		}
		buf, err := json.Marshal(subscription)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("[OnOpen] Connected, subscribe: ", string(buf))
		client.Send(string(buf))
	}

	client.OnMessage = func(msg string) {
		log.Println("[OnMessage]", msg)
	}

	client.OnClose = func() {
		log.Println("[OnClose] Disconnected.")
	}

	client.OnError = func(err error) {
		log.Println("[OnError]", err)
	}

	if err := client.Connect(); err != nil {
		log.Fatal("Failed to connect:", err)
	}

	// Wait for Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Interrupt received. Shutting down...")
	client.Close()
}
