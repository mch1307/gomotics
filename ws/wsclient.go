package ws

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mch1307/gomotics/log"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Add CheckOrigin forced to true to avoid "Origin header value not allowed"
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSClient handles the ws client conn.
// use send channel as input for pushing msg to client
type WSClient struct {
	hub *ClientPool
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel for outbound messages.
	send chan []byte
}

// ServeWebSocket called by the http endpoint. Calls serveWS function and adds the WSPool
// to original connection
func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	serveWS(WSPool, w, r)
}

// ServeWS upgrade connection to WS, link it to the WSPool
// and start 1 read and 1 write goroutine per connection
func serveWS(hub *ClientPool, w http.ResponseWriter, r *http.Request) {
	log.Debug("serveWS started")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	client := &WSClient{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	log.Debug("serveWS starting goroutines")
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

// readPump goroutine to read messages from the websocket connection to the WSClient.
// Mainly for control msg
func (c *WSClient) readPump() {
	log.Debug("started")
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Debug("error: %v", err)
			}
			break
		}
	}
}

// writePump goroutine pushes messages from the ClientPool
// to the websocket connection. One goroutine per connection
func (c *WSClient) writePump() {
	log.Debug("started")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// Add queued events to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				//w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
