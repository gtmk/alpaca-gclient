package alpacaio

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MaxConnectionAttempts = 3
)

var (
	once   sync.Once
	stream *Stream
)

type credentials struct {
	apiKey         string
	secretKey      string
	streamEndpoint string
}

type Stream struct {
	sync.Mutex
	sync.Once
	conn                  *websocket.Conn
	authenticated, closed atomic.Value
	credentials           credentials

	MessageC chan []byte
	ErrorC   chan error
}

func GetStream(apiKey, secretKey, streamEndpoint string) (*Stream, error) {
	once.Do(func() {
		stream = &Stream{
			authenticated: atomic.Value{},
			MessageC:      make(chan []byte, 100),
			ErrorC:        make(chan error, 100),
			credentials: credentials{apiKey: apiKey,
				secretKey:      secretKey,
				streamEndpoint: streamEndpoint},
		}
		stream.authenticated.Store(false)
		stream.closed.Store(false)
	})
	err := stream.register()
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (s *Stream) register() error {
	var err error
	if s.conn == nil {
		s.conn, err = s.openSocket()
		if err != nil {
			return err
		}
	}
	if err = s.auth(); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Subscribe(channel string) error {
	s.Do(func() {
		go s.start()
	})
	if err := s.sub(channel); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Unsubscribe(channel string) error {
	var err error
	if s.conn == nil {
		return errors.New("connection has not been initialized")
	}
	if err = s.auth(); err != nil {
		return err
	}
	if err = s.unsub(channel); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Close() error {
	s.Lock()
	defer s.Unlock()

	if s.conn == nil {
		return nil
	}

	if err := s.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); err != nil {
		return err
	}
	s.closed.Store(true)
	return s.conn.Close()
}

func (s *Stream) openSocket() (*websocket.Conn, error) {
	connectionAttempts := 0
	for connectionAttempts < MaxConnectionAttempts {
		c, _, err := websocket.DefaultDialer.Dial(s.credentials.streamEndpoint, nil)
		if err == nil {
			return c, nil
		}
		if connectionAttempts == MaxConnectionAttempts {
			return nil, err
		}
		time.Sleep(1 * time.Second)
		connectionAttempts++
	}
	return nil, fmt.Errorf("Error: Could not open Alpaca stream (max retries exceeded).")
}

func (s *Stream) isAuthenticated() bool {
	return s.authenticated.Load().(bool)
}

func (s *Stream) auth() error {
	s.Lock()
	defer s.Unlock()

	if s.isAuthenticated() {
		return nil
	}
	authRequest := ClientMsg{
		Action: "authenticate",
		Data: map[string]interface{}{
			"key_id":     s.credentials.apiKey,
			"secret_key": s.credentials.secretKey,
		},
	}

	if err := s.conn.WriteJSON(authRequest); err != nil {
		return err
	}
	msg := ServerMsg{}
	// ensure the auth response comes in a timely manner
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	defer s.conn.SetReadDeadline(time.Time{})

	if err := s.conn.ReadJSON(&msg); err != nil {
		return err
	}
	m := msg.Data.(map[string]interface{})

	if !strings.EqualFold(m["status"].(string), "authorized") {
		return fmt.Errorf("failed to authorize alpaca stream")
	}
	s.authenticated.Store(true)
	return nil
}

func (s *Stream) start() {
	for {
		code, bts, err := s.conn.ReadMessage()
		if err != nil {
			if s.closed.Load().(bool) {
				return
			} else if websocket.IsCloseError(err, code) {
				err := s.reconnect()
				if err != nil {
					s.ErrorC <- err
					return
				}
				continue
			} else {
				s.ErrorC <- err
				continue
			}
		}
		s.MessageC <- bts
	}
}

func (s *Stream) sub(channel string) error {
	s.Lock()
	defer s.Unlock()

	subReq := ClientMsg{
		Action: "listen",
		Data: map[string]interface{}{
			"streams": []interface{}{
				channel,
			},
		},
	}
	if err := s.conn.WriteJSON(subReq); err != nil {
		return err
	}
	return nil
}

func (s *Stream) unsub(channel string) error {
	s.Lock()
	defer s.Unlock()
	subReq := ClientMsg{
		Action: "unlisten",
		Data: map[string]interface{}{
			"streams": []interface{}{
				channel,
			},
		},
	}
	err := s.conn.WriteJSON(subReq)
	return err
}

func (s *Stream) reconnect() error {
	var err error
	s.authenticated.Store(false)
	if s.conn, err = s.openSocket(); err != nil {
		return err
	}
	if err = s.auth(); err != nil {
		return err
	}
	return nil
}

type ClientMsg struct {
	Action string      `json:"action" msgpack:"action"`
	Data   interface{} `json:"data" msgpack:"data"`
}

type ServerMsg struct {
	Stream string      `json:"stream" msgpack:"stream"`
	Data   interface{} `json:"data"`
}
