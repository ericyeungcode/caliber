package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/ericyeungcode/caliber"
)

const (
	SubTypeSubscribe = "subscribe"
)

type SubReq struct {
	Type        string   `json:"type"`
	Instruments []string `json:"instruments"`
	Channels    []string `json:"channels"`
	Interval    string   `json:"interval"`
}

type PrivateSubSeq struct {
	*SubReq
	Token string `json:"token"`
}

func main() {
	wsURL := "wss://betaws.bitexch.dev"

	client := caliber.NewWSClient(wsURL, &tls.Config{
		InsecureSkipVerify: true, // ⚠️ only for dev / testing
	})

	client.OnOpen = func() {
		var subscription = &SubReq{
			Type:        SubTypeSubscribe,
			Instruments: []string{"BTC-USDT-PERPETUAL"},
			Channels:    []string{"order_book.10.10"},
			Interval:    "raw",
		}
		buf, err := json.Marshal(&subscription)
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

	// go func() {
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		client.Send("ping from client")
	// 	}
	// }()

	// Wait for Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Interrupt received. Shutting down...")
	client.Close()
}
