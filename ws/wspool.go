package ws

import "github.com/mch1307/gomotics/log"

// WSPool instance to hold ClientPool
var WSPool *ClientPool

// Initialize instantiate a ClientPool at startup (through init)
// the start the Run process
func Initialize() {
	WSPool = newClientPool()
	go WSPool.Run()
}

// ClientPool maintains the set of active clients and broadcasts messages to the
// clients. Use the Broadcast channel to push msg to all clients
type ClientPool struct {
	// Clients: Registered clients.
	clients map[*WSClient]bool
	// Broadcast Inbound messages from the clients.
	Broadcast chan []byte
	// Register requests from the clients.
	register chan *WSClient
	// Unregister requests from clients.
	Unregister chan *WSClient
}

// newClientPool instantiates the ClientPool
func newClientPool() *ClientPool {
	log.Debug("new ws pool created")
	return &ClientPool{
		Broadcast:  make(chan []byte),
		register:   make(chan *WSClient),
		Unregister: make(chan *WSClient),
		clients:    make(map[*WSClient]bool),
	}
}

// Run activates the ClientPool so that we can manage
// client connections
func (cp *ClientPool) Run() {
	for {
		select {
		case client := <-cp.register:
			cp.clients[client] = true
		case client := <-cp.Unregister:
			if _, ok := cp.clients[client]; ok {
				delete(cp.clients, client)
				close(client.send)
			}
		case message := <-cp.Broadcast:
			for client := range cp.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(cp.clients, client)
				}
			}
		}
	}
}
