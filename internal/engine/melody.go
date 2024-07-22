package engine

import (
	"github.com/gorilla/websocket"
	"github.com/olahol/melody"
	"net"
	"net/http"
)

type Symphony struct {
	*melody.Melody
	sessions map[*melody.Session]*Chession
}

func NewSymphony(m *melody.Melody) *Symphony {
	return &Symphony{Melody: m, sessions: make(map[*melody.Session]*Chession)}
}

func (s *Symphony) HandleConnect(fn func(SessionInterface)) {
	s.Melody.HandleConnect(func(session *melody.Session) {
		s.sessions[session] = NewChession(session)
		fn(s.sessions[session])
	})
}

func (s *Symphony) HandleDisconnect(fn func(SessionInterface)) {
	s.Melody.HandleDisconnect(func(session *melody.Session) {
		fn(s.sessions[session])
		delete(s.sessions, session)
	})
}

func (s *Symphony) HandlePong(fn func(SessionInterface)) {
	s.Melody.HandlePong(func(session *melody.Session) {
		fn(s.sessions[session])
	})
}

func (s *Symphony) HandleMessage(fn func(SessionInterface, []byte)) {
	s.Melody.HandleMessage(func(session *melody.Session, msg []byte) {
		fn(s.sessions[session], msg)
	})
}

func (s *Symphony) HandleMessageBinary(fn func(SessionInterface, []byte)) {
	s.Melody.HandleMessageBinary(func(session *melody.Session, msg []byte) {
		fn(s.sessions[session], msg)
	})
}

func (s *Symphony) HandleSentMessage(fn func(SessionInterface, []byte)) {
	s.Melody.HandleSentMessage(func(session *melody.Session, msg []byte) {
		fn(s.sessions[session], msg)
	})
}

func (s *Symphony) HandleSentMessageBinary(fn func(SessionInterface, []byte)) {
	s.Melody.HandleSentMessageBinary(func(session *melody.Session, msg []byte) {
		fn(s.sessions[session], msg)
	})
}

func (s *Symphony) HandleError(fn func(SessionInterface, error)) {
	s.Melody.HandleError(func(session *melody.Session, err error) {
		fn(s.sessions[session], err)
	})
}

func (s *Symphony) HandleClose(fn func(SessionInterface, int, string) error) {
	s.Melody.HandleClose(func(session *melody.Session, code int, reason string) error {
		return fn(s.sessions[session], code, reason)
	})
}

func (s *Symphony) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return s.Melody.HandleRequest(w, r)
}

func (s *Symphony) HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]any) error {
	return s.Melody.HandleRequestWithKeys(w, r, keys)
}

func (s *Symphony) Broadcast(msg []byte) error {
	return s.Melody.Broadcast(msg)
}

func (s *Symphony) Sessions() ([]SessionInterface, error) {
	var sessions []SessionInterface
	for _, session := range s.sessions {
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (s *Symphony) Close() error {
	return s.Melody.Close()
}

func (s *Symphony) CloseWithMsg(msg []byte) error {
	return s.Melody.CloseWithMsg(msg)
}

func (s *Symphony) Len() int {
	return s.Melody.Len()
}

func (s *Symphony) IsClosed() bool {
	return s.Melody.IsClosed()
}

type MelodyInterface interface {
	HandleConnect(fn func(SessionInterface))
	HandleDisconnect(fn func(SessionInterface))
	HandlePong(fn func(SessionInterface))
	HandleMessage(fn func(SessionInterface, []byte))
	HandleMessageBinary(fn func(SessionInterface, []byte))
	HandleSentMessage(fn func(SessionInterface, []byte))
	HandleSentMessageBinary(fn func(SessionInterface, []byte))
	HandleError(fn func(SessionInterface, error))
	HandleClose(fn func(SessionInterface, int, string) error)
	HandleRequest(w http.ResponseWriter, r *http.Request) error
	HandleRequestWithKeys(w http.ResponseWriter, r *http.Request, keys map[string]any) error
	Broadcast(msg []byte) error
	Sessions() ([]SessionInterface, error)
	Close() error
	CloseWithMsg(msg []byte) error
	Len() int
	IsClosed() bool
}

type Chession struct {
	*melody.Session
}

func NewChession(session *melody.Session) *Chession {
	return &Chession{Session: session}
}

func (c Chession) Write(msg []byte) error {
	return c.Session.Write(msg)
}

func (c Chession) WriteBinary(msg []byte) error {
	return c.Session.WriteBinary(msg)
}

func (c Chession) Close() error {
	return c.Session.Close()
}

func (c Chession) CloseWithMsg(msg []byte) error {
	return c.Session.CloseWithMsg(msg)
}

func (c Chession) Set(key string, value any) {
	c.Session.Set(key, value)
}

func (c Chession) Get(key string) (value any, exists bool) {
	return c.Session.Get(key)
}

func (c Chession) MustGet(key string) any {
	return c.Session.MustGet(key)
}

func (c Chession) UnSet(key string) {
	c.Session.UnSet(key)
}

func (c Chession) IsClosed() bool {
	return c.Session.IsClosed()
}

func (c Chession) LocalAddr() net.Addr {
	return c.Session.LocalAddr()
}

func (c Chession) RemoteAddr() net.Addr {
	return c.Session.RemoteAddr()
}

func (c Chession) WebsocketConnection() *websocket.Conn {
	return c.Session.WebsocketConnection()
}

type SessionInterface interface {
	Write(msg []byte) error
	WriteBinary(msg []byte) error
	Close() error
	CloseWithMsg(msg []byte) error
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	MustGet(key string) any
	UnSet(key string)
	IsClosed() bool
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	WebsocketConnection() *websocket.Conn
}
