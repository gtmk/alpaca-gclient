// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package alpacaio

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(in *jlexer.Lexer, out *StreamingServerMsges) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StreamingServerMsges, 0, 2)
			} else {
				*out = StreamingServerMsges{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 StreamingServerMsges
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(out *jwriter.Writer, in StreamingServerMsges) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StreamingServerMsges) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamingServerMsges) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamingServerMsges) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamingServerMsges) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(in *jlexer.Lexer, out *StreamingServerMsg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "i":
			out.TradeID = int64(in.Int64())
		case "x":
			out.Exchange = string(in.String())
		case "p":
			out.Price = float32(in.Float32())
		case "s":
			out.Size = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "c":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]string, 0, 4)
					} else {
						out.Conditions = []string{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Conditions = append(out.Conditions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "z":
			out.Tape = string(in.String())
		case "bx":
			out.BidExchange = int32(in.Int32())
		case "ax":
			out.AskExchange = int32(in.Int32())
		case "bp":
			out.BidPrice = float32(in.Float32())
		case "ap":
			out.AskPrice = float32(in.Float32())
		case "bs":
			out.BidSize = int32(in.Int32())
		case "as":
			out.AskSize = int32(in.Int32())
		case "v":
			out.Volume = int32(in.Int32())
		case "o":
			out.Open = float32(in.Float32())
		case "h":
			out.High = float32(in.Float32())
		case "l":
			out.Low = float32(in.Float32())
		case "c":
			out.Close = float32(in.Float32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(out *jwriter.Writer, in StreamingServerMsg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.Int64(int64(in.TradeID))
	}
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix)
		out.String(string(in.Exchange))
	}
	{
		const prefix string = ",\"p\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Int32(int32(in.Size))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Conditions {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	{
		const prefix string = ",\"bx\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidExchange))
	}
	{
		const prefix string = ",\"ax\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskExchange))
	}
	{
		const prefix string = ",\"bp\":"
		out.RawString(prefix)
		out.Float32(float32(in.BidPrice))
	}
	{
		const prefix string = ",\"ap\":"
		out.RawString(prefix)
		out.Float32(float32(in.AskPrice))
	}
	{
		const prefix string = ",\"bs\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidSize))
	}
	{
		const prefix string = ",\"as\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskSize))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.Int32(int32(in.Volume))
	}
	{
		const prefix string = ",\"o\":"
		out.RawString(prefix)
		out.Float32(float32(in.Open))
	}
	{
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Float32(float32(in.High))
	}
	{
		const prefix string = ",\"l\":"
		out.RawString(prefix)
		out.Float32(float32(in.Low))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Float32(float32(in.Close))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StreamingServerMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamingServerMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamingServerMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamingServerMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(in *jlexer.Lexer, out *StreamTrades) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StreamTrades, 0, 0)
			} else {
				*out = StreamTrades{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 StreamTrade
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(out *jwriter.Writer, in StreamTrades) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StreamTrades) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamTrades) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamTrades) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamTrades) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(in *jlexer.Lexer, out *StreamTrade) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "i":
			out.TradeID = int64(in.Int64())
		case "x":
			out.Exchange = string(in.String())
		case "p":
			out.Price = float32(in.Float32())
		case "s":
			out.Size = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "c":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]string, 0, 4)
					} else {
						out.Conditions = []string{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Conditions = append(out.Conditions, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "z":
			out.Tape = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(out *jwriter.Writer, in StreamTrade) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.Int64(int64(in.TradeID))
	}
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix)
		out.String(string(in.Exchange))
	}
	{
		const prefix string = ",\"p\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Int32(int32(in.Size))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Conditions {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StreamTrade) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamTrade) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamTrade) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamTrade) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(in *jlexer.Lexer, out *StreamQuotes) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StreamQuotes, 0, 0)
			} else {
				*out = StreamQuotes{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 StreamQuote
			(v13).UnmarshalEasyJSON(in)
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(out *jwriter.Writer, in StreamQuotes) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			(v15).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StreamQuotes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamQuotes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamQuotes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamQuotes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(in *jlexer.Lexer, out *StreamQuote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "c":
			out.Condition = int32(in.Int32())
		case "bx":
			out.BidExchange = int32(in.Int32())
		case "ax":
			out.AskExchange = int32(in.Int32())
		case "bp":
			out.BidPrice = float32(in.Float32())
		case "ap":
			out.AskPrice = float32(in.Float32())
		case "bs":
			out.BidSize = int32(in.Int32())
		case "as":
			out.AskSize = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "z":
			out.Tape = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(out *jwriter.Writer, in StreamQuote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Int32(int32(in.Condition))
	}
	{
		const prefix string = ",\"bx\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidExchange))
	}
	{
		const prefix string = ",\"ax\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskExchange))
	}
	{
		const prefix string = ",\"bp\":"
		out.RawString(prefix)
		out.Float32(float32(in.BidPrice))
	}
	{
		const prefix string = ",\"ap\":"
		out.RawString(prefix)
		out.Float32(float32(in.AskPrice))
	}
	{
		const prefix string = ",\"bs\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidSize))
	}
	{
		const prefix string = ",\"as\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskSize))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StreamQuote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamQuote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamQuote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamQuote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(in *jlexer.Lexer, out *StreamAggregates) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StreamAggregates, 0, 0)
			} else {
				*out = StreamAggregates{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v16 StreamAggregate
			(v16).UnmarshalEasyJSON(in)
			*out = append(*out, v16)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(out *jwriter.Writer, in StreamAggregates) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v17, v18 := range in {
			if v17 > 0 {
				out.RawByte(',')
			}
			(v18).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StreamAggregates) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamAggregates) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamAggregates) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamAggregates) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(in *jlexer.Lexer, out *StreamAggregate) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "v":
			out.Volume = int32(in.Int32())
		case "o":
			out.Open = float32(in.Float32())
		case "h":
			out.High = float32(in.Float32())
		case "l":
			out.Low = float32(in.Float32())
		case "c":
			out.Close = float32(in.Float32())
		case "s":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(out *jwriter.Writer, in StreamAggregate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.Int32(int32(in.Volume))
	}
	{
		const prefix string = ",\"o\":"
		out.RawString(prefix)
		out.Float32(float32(in.Open))
	}
	{
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Float32(float32(in.High))
	}
	{
		const prefix string = ",\"l\":"
		out.RawString(prefix)
		out.Float32(float32(in.Low))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Float32(float32(in.Close))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StreamAggregate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamAggregate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamAggregate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamAggregate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(l, v)
}