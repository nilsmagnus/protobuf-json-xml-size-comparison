package main

import (
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nilsmagnus/grpc-samples/sample"
)

func main() {

	show := flag.String("show", "nothing", "what to show: json, xml, proto, protolen")
	numberOfMapEntries := flag.Int("entries", 100, "How many entries in map")

	flag.Parse()

	tickers := make(map[string]float32)

	for i := 1; i <= *numberOfMapEntries; i++ {
		tickers[fmt.Sprintf("ticker-%d", i)] = 3.2
	}

	protosome := &sample.Test{
		Query:         "myQuery",
		PageNumber:    42,
		ResultPerPage: 100,
		TickerValues:  tickers,
	}

	switch *show {
	case "protolen":
		data, _ := proto.Marshal(protosome)

		fmt.Printf("size of protoencoded : %d\n", binary.Size(data))
	case "proto":
		data, _ := proto.Marshal(protosome)
		fmt.Println(data)
	case "json":
		jsonified, _ := json.Marshal(protosome)
		fmt.Println(string(jsonified))
	case "jsonlen":
		jsonified, _ := json.Marshal(protosome)
		fmt.Printf("json length(raw): %s", len(jsonified))
	case "xml":
		xmlified, _ := xml.Marshal(protosome)
		fmt.Println(string(xmlified))
	case "xmllen":
		xmlified, _ := xml.Marshal(protosome)
		fmt.Printf("xml lenght(raw):", len(xmlified))

	default:
		fmt.Printf("unknown show-option: %s\n", *show)
	}

}
