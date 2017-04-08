package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nilsmagnus/grpc-samples/sample"
	"math/rand"
	"time"
)

func main() {

	show := flag.String("show", "nothing", "what to show: json, jsonlen, proto, protolen")
	numberOfMapEntries := flag.Int("entries", 100, "How many entries in map")

	flag.Parse()

	tickers := make([]*sample.Ticker, *numberOfMapEntries)

	for i := 0; i < *numberOfMapEntries; i++ {
		tickers[i] = &sample.Ticker{
			Value: rand.Float32() * 10.0,
			Name:  RandStringRunes(3),
		}
	}

	protosome := &sample.Test{
		Query:         "myQuery",
		PageNumber:    42,
		ResultPerPage: 100,
		Tickers:       tickers,
	}

	switch *show {
	case "protolen":
		data, _ := proto.Marshal(protosome)
		fmt.Printf("size of protoencoded : %d\n", len(data))
	case "proto":
		data, _ := proto.Marshal(protosome)
		fmt.Println(data)
	case "json":
		jsonified, _ := json.Marshal(protosome)
		fmt.Println(string(jsonified))
	case "jsonlen":
		jsonified, _ := json.Marshal(protosome)
		fmt.Printf("json length(raw): %d\n", len(jsonified))
	default:
		fmt.Printf("unknown show-option: %s\n", *show)
	}

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
