package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	clients    map[*websocket.Conn]bool
	lock       sync.Mutex
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast  chan []byte
}

func NewManager() *Manager {
	return &Manager{
		clients:    make(map[*websocket.Conn]bool),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		broadcast:  make(chan []byte),
	}
}

func (m *Manager) Run() {
	for {
		select {
		case conn := <-m.register:
			m.lock.Lock()
			m.clients[conn] = true
			m.lock.Unlock()
		case conn := <-m.unregister:
			m.lock.Lock()
			if _, ok := m.clients[conn]; ok {
				delete(m.clients, conn)
				conn.Close()
			}
			m.lock.Unlock()
		case msg := <-m.broadcast:
			m.lock.Lock()
			for c := range m.clients {
				_ = c.WriteMessage(websocket.TextMessage, msg)
			}
			m.lock.Unlock()
		}
	}
}

func (m *Manager) Register(conn *websocket.Conn) {
	m.register <- conn
}

func (m *Manager) Unregister(conn *websocket.Conn) {
	m.unregister <- conn
}

func (m *Manager) BroadcastMessage(msg []byte) {
	m.broadcast <- msg
}
