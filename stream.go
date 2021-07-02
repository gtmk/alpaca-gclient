package alpacaio

import (
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

	channels Channels
}

type Channels struct {
	TradeChannel   chan StreamingServerMsg
	QuoteChannel   chan StreamingServerMsg
	BarChannel     chan StreamingServerMsg
	ErrorChannel   chan StreamingServerMsg
	DefaultChannel chan StreamingServerMsg
}

func (cs Channels) close() {
	close(cs.TradeChannel)
	close(cs.QuoteChannel)
	close(cs.BarChannel)
	close(cs.ErrorChannel)
	close(cs.DefaultChannel)
}

func GetStream(apiKey, secretKey, streamEndpoint string) (Channels, error) {
	once.Do(func() {
		stream = &Stream{
			authenticated: atomic.Value{},
			channels: Channels{
				TradeChannel:   make(chan StreamingServerMsg, 100),
				QuoteChannel:   make(chan StreamingServerMsg, 100),
				ErrorChannel:   make(chan StreamingServerMsg, 100),
				BarChannel:     make(chan StreamingServerMsg, 100),
				DefaultChannel: make(chan StreamingServerMsg, 100),
			},
			credentials: credentials{apiKey: apiKey,
				secretKey:      secretKey,
				streamEndpoint: streamEndpoint},
		}
		stream.authenticated.Store(false)
		stream.closed.Store(false)
	})
	err := stream.register()
	if err != nil {
		return Channels{}, err
	}
	return stream.channels, nil
}

func (s *Stream) register() error {
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

func Subscribe(channel string, tickers []string) error {
	if err := stream.sub(channel, tickers); err != nil {
		return err
	}
	return nil
}

func Unsubscribe(channel string, tickers []string) error {
	var err error
	if stream.conn == nil {
		return errors.New("connection has not been initialized")
	}
	if err = stream.auth(); err != nil {
		return err
	}
	if err = stream.unsub(channel, tickers); err != nil {
		return err
	}
	return nil
}

func Close() error {
	stream.Lock()
	defer stream.Unlock()

	if stream.conn == nil {
		return nil
	}

	if err := stream.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); err != nil {
		return err
	}
	stream.closed.Store(true)
	stream.channels.close()
	return stream.conn.Close()
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
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	defer s.conn.SetReadDeadline(time.Time{})

	if err := s.conn.ReadJSON(&connected); err != nil {
		return err
	}
	if err := s.conn.ReadJSON(&authorized); err != nil {
		return err
	}

	for _, m := range append(connected, authorized...) {
		if m.Message == "authenticated" {
			s.authenticated.Store(true)
			return nil
		}
	}
	return fmt.Errorf("failed to authorize alpaca stream")
}

func (s *Stream) errorMake(msg string) StreamingServerMsg {
	return StreamingServerMsg{
		Event:   `error`,
		Code:    10000,
		Message: msg,
	}
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
					s.channels.ErrorChannel <- s.errorMake(err.Error())
				}
				continue
			} else {
				s.channels.ErrorChannel <- s.errorMake(err.Error())
				continue
			}
		}
		go func() {
			if err := s.parseStreamMessages(bts); err != nil {
				s.channels.ErrorChannel <- s.errorMake(err.Error())
			}
		}()
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
		subReq.Quotes = tickers
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

func (s *Stream) parseStreamMessages(bts []byte) error {
	var msgTypes StreamingServerMsges
	err := ej.Unmarshal(bts, &msgTypes)
	if err != nil {
		return err
	}
	for _, msg := range msgTypes {
		switch msg.Event {
		case `t`:
			s.channels.TradeChannel <- msg
		case `q`:
			s.channels.QuoteChannel <- msg
		case `b`:
			s.channels.BarChannel <- msg
		case `error`:
			s.channels.ErrorChannel <- msg
		default:
			s.channels.DefaultChannel <- msg
		}
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
