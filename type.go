package alpacaio

import "time"

type StreamingServerMsg struct {
	Event      string    `json:"T"`
	Symbol     string    `json:"S"`
	TradeID    int64     `json:"i"`
	Exchange   string    `json:"x"`
	Price      float32   `json:"p"`
	Size       int32     `json:"s"`
	Timestamp  time.Time `json:"t"`
	Conditions []string  `json:"c"`
	Tape       string    `json:"z"`
	// Quote
	BidExchange int32   `json:"bx"`
	AskExchange int32   `json:"ax"`
	BidPrice    float32 `json:"bp"`
	AskPrice    float32 `json:"ap"`
	BidSize     int32   `json:"bs"`
	AskSize     int32   `json:"as"`
	// Bar
	Volume int32   `json:"v"`
	Open   float32 `json:"o"`
	High   float32 `json:"h"`
	Low    float32 `json:"l"`
	Close  float32 `json:"c"`
}

//easyjson:json
type StreamingServerMsges []StreamingServerMsges

// StreamTrade is the structure that defines a trade that
// polygon transmits via websocket protocol.
type StreamTrade struct {
	Event      string    `json:"T"`
	Symbol     string    `json:"S"`
	TradeID    int64     `json:"i"`
	Exchange   string    `json:"x"`
	Price      float32   `json:"p"`
	Size       int32     `json:"s"`
	Timestamp  time.Time `json:"t"`
	Conditions []string  `json:"c"`
	Tape       string    `json:"z"`
}

//easyjson:json
type StreamTrades []StreamTrade

// StreamQuote is the structure that defines a quote that
// polygon transmits via websocket protocol.
type StreamQuote struct {
	Event       string    `json:"T"`
	Symbol      string    `json:"S"`
	Condition   int32     `json:"c"`
	BidExchange int32     `json:"bx"`
	AskExchange int32     `json:"ax"`
	BidPrice    float32   `json:"bp"`
	AskPrice    float32   `json:"ap"`
	BidSize     int32     `json:"bs"`
	AskSize     int32     `json:"as"`
	Timestamp   time.Time `json:"t"`
	Tape        string    `json:"z"`
}

//easyjson:json
type StreamQuotes []StreamQuote

// StreamAggregate is the structure that defines an aggregate that
// polygon transmits via websocket protocol.
type StreamAggregate struct {
	Event     string    `json:"T"`
	Symbol    string    `json:"S"`
	Volume    int32     `json:"v"`
	Open      float32   `json:"o"`
	High      float32   `json:"h"`
	Low       float32   `json:"l"`
	Close     float32   `json:"c"`
	Timestamp time.Time `json:"s"`
}

//easyjson:json
type StreamAggregates []StreamAggregate
