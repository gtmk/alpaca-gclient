package alpacaio

import (
	json "encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	ej "github.com/mailru/easyjson"
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

func (s *Stream) Register() error {
	var err error
	if s.conn == nil {
		s.conn, err = s.openSocket()
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to open with err %+v", err))
			return err
		}
	}
	if err = s.auth(); err != nil {
		return err
	}
	go s.start()
	return nil
}

func (s *Stream) Subscribe(channel string, tickers []string) error {
	if err := s.sub(channel, tickers); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Unsubscribe(channel string, tickers []string) error {
	var err error
	if s.conn == nil {
		return errors.New("connection has not been initialized")
	}
	if err = s.auth(); err != nil {
		return err
	}
	if err = s.unsub(channel, tickers); err != nil {
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
	close(s.MessageC)
	close(s.ErrorC)
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
	authRequest := ClientAuthMsg{
		Action:    "auth",
		APIKey:    s.credentials.apiKey,
		SecretKey: s.credentials.secretKey,
	}

	if err := s.conn.WriteJSON(authRequest); err != nil {
		return err
	}
	connected := []ServerAuthMsg{}
	authorized := []ServerAuthMsg{}
	// ensure the auth response comes in a timely manner
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	defer s.conn.SetReadDeadline(time.Time{})

	if err := s.conn.ReadJSON(&connected); err != nil {
		return err
	}
	if err := s.conn.ReadJSON(&authorized); err != nil {
		return err
	}

	fmt.Println(append(connected, authorized...))
	for _, m := range append(connected, authorized...) {
		if m.Message == "authenticated" {
			s.authenticated.Store(true)
			return nil
		}
	}
	return fmt.Errorf("failed to authorize alpaca stream")
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
					s.conn = nil
					fmt.Println(fmt.Sprintf("unknown error %+v", err))
					return
				}
				continue
			} else {
				s.ErrorC <- err
				s.conn = nil
				fmt.Println(fmt.Sprintf("unknown error %+v", err))
				return
			}
		}
		s.MessageC <- bts
	}
}

func (s *Stream) sub(channel string, tickers []string) error {
	s.Lock()
	defer s.Unlock()

	subs := make(map[string]interface{}, len(tickers))
	for _, t := range tickers {
		subs[channel] = t
	}
	subReq := ClientSubcribeMsg{
		Action: "subscribe",
	}
	switch channel {
	case `trades`:
		subReq.Trades = tickers
	case `quotes`:
		subReq.Trades = tickers
	case `bars`:
		subReq.Bars = tickers
	default:
		return fmt.Errorf("unknown channel")
	}
	if err := s.conn.WriteJSON(subReq); err != nil {
		return err
	}
	return nil
}

func (s *Stream) unsub(channel string, tickers []string) error {
	s.Lock()

	defer s.Unlock()
	subReq := ClientSubcribeMsg{
		Action: "unsubscribe",
	}
	switch channel {
	case `trades`:
		subReq.Trades = tickers
	case `quotes`:
		subReq.Trades = tickers
	case `bars`:
		subReq.Bars = tickers
	default:
		return fmt.Errorf("unknown channel")
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

type ClientAuthMsg struct {
	Action    string `json:"action,omitempty" msgpack:"action"`
	APIKey    string `json:"key,omitempty"`
	SecretKey string `json:"secret",omitempty`
}

type ClientSubcribeMsg struct {
	Action string   `json:"action,omitempty" msgpack:"action"`
	Trades []string `json:"trades,omitempty"`
	Quotes []string `json:"quotes,omitempty"`
	Bars   []string `json:"bars,omitempty"`
}

type ServerAuthMsg struct {
	EventName string `json:"T" msgpack:"stream"`
	Message   string `json:"msg"`
}

func ParseEvents(bts []byte, isEJ ...bool) (StreamingServerMsges, error) {
	iEJ := true
	var out StreamingServerMsges
	var err error
	if iEJ {
		err = ej.Unmarshal(bts, &out)
	} else {
		err = json.Unmarshal(bts, &out)
	}
	return out, err
}
