package caliber

import (
	"crypto/tls"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
)

const (
	maxRetryDelay  = 10 * time.Second
	initialBackoff = 1 * time.Second
)

type WSClient struct {
	wsHost           string
	tlsConfig        *tls.Config
	conn             *websocket.Conn
	sendQueue        chan string
	quit             chan struct{}
	closed           atomic.Bool
	reconnectingFlag atomic.Int32

	OnOpen    func()
	OnMessage func(string)
	OnClose   func()
	OnError   func(error)
}

/*

client := wsclient.NewWSClient(wsURL, &tls.Config{
	InsecureSkipVerify: true, // ⚠️ only for dev / testing
})

*/

// NewWSClient constructs a new client
func NewWSClient(wsHost string, tlsConfig *tls.Config) *WSClient {
	return &WSClient{
		wsHost:    wsHost,
		tlsConfig: tlsConfig,
		sendQueue: make(chan string, 1024),
		quit:      make(chan struct{}),
	}
}

// Connect establishes the WebSocket connection
func (c *WSClient) Connect() error {
	log.Println("Connecting to", c.wsHost)

	u, err := url.Parse(c.wsHost)
	if err != nil {
		return err
	}

	dialer := websocket.DefaultDialer
	if u.Scheme == "wss" && c.tlsConfig != nil {
		dialer = &websocket.Dialer{
			TLSClientConfig: c.tlsConfig,
		}
	}

	conn, _, err := dialer.Dial(c.wsHost, nil)
	if err != nil {
		if c.OnError != nil {
			c.OnError(err)
		}
		return err
	}

	c.conn = conn
	if c.OnOpen != nil {
		c.OnOpen()
	}

	go c.readLoop()
	go c.writeLoop()

	return nil
}

func (c *WSClient) readLoop() {
	for {
		select {
		case <-c.quit:
			log.Println("Stopping readLoop")
			return
		default:
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				c.reconnect()
				return
			}
			if c.OnMessage != nil {
				c.OnMessage(string(msg))
			}
		}
	}
}

func (c *WSClient) writeLoop() {
	for {
		select {
		case <-c.quit:
			log.Println("Stopping writeLoop")
			return
		case msg := <-c.sendQueue:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
				c.reconnect()
				return
			}
		}
	}
}

func (c *WSClient) Send(msg string) {
	select {
	case c.sendQueue <- msg:
	default:
		log.Println("Send queue full, dropping message")
	}
}

func (c *WSClient) reconnect() {
	if !c.reconnectingFlag.CompareAndSwap(0, 1) {
		log.Println("reconnect() is working...")
		return
	}
	defer c.reconnectingFlag.Store(0)

	if c.conn != nil {
		c.conn.Close()
	}
	if c.OnClose != nil {
		c.OnClose()
	}

	backoff := initialBackoff
	for {
		if c.closed.Load() {
			return
		}
		log.Println("Attempting to reconnect...")
		err := c.Connect()
		if err == nil {
			log.Println("Reconnected successfully")
			break
		}
		if c.OnError != nil {
			c.OnError(err)
		}
		time.Sleep(backoff)
		if backoff < maxRetryDelay {
			backoff *= 2
		} else {
			backoff = maxRetryDelay
		}
	}
}

// Close shuts down the client cleanly
func (c *WSClient) Close() {
	log.Println("Closing connection")
	c.closed.Store(true)
	close(c.quit)
	if c.conn != nil {
		c.conn.Close()
	}
	if c.OnClose != nil {
		c.OnClose()
	}
}
