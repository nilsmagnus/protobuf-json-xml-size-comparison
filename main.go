package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nilsmagnus/grpc-samples/sample"
	"math/rand"
	"time"
	"bytes"
	"compress/gzip"
)

func main() {
	for _, dataSize := range []int{0, 1, 2, 10, 20, 200, 2000, 20000} {
		protoStruct := createTestDatata(dataSize)
		jsonl, gzJsonlen,  protol := jsonProtoLengts(protoStruct)
		fmt.Printf("# %d tickers, json: %d, gzJson %d, proto: %d , proto vs json %f ,proto vs gzipjson %f \n", dataSize, jsonl, gzJsonlen, protol, float32(protol)/float32(jsonl), float32(protol)/float32(gzJsonlen))
		fmt.Printf("| %d |  %d | %d | %d | %f | %f | \n", dataSize, jsonl, gzJsonlen, protol, float32(protol)/float32(jsonl), float32(protol)/float32(gzJsonlen))
	}

}
func createTestDatata(numberOfEntries int) *sample.Test {

	tickers := make([]*sample.Ticker, numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		tickers[i] = &sample.Ticker{
			Value: rand.Float32() * 10.0,
			Name:  RandStringRunes(3),
		}
	}

	return &sample.Test{
		Query:         "myQuery",
		PageNumber:    42,
		ResultPerPage: 100,
		Tickers:       tickers,
	}

}

func jsonProtoLengts(protoSome *sample.Test) (jsonLen, gzipJsonLen, protoLen int) {
	data, _ := proto.Marshal(protoSome)
	protoLen = len(data)
	jsonified, _ := json.Marshal(protoSome)
	jsonLen = len(jsonified)
	gzipJsonLen = gzipLen(jsonified)
	return
}

func gzipLen(jsonData []byte) int {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(jsonData); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}

	return b.Len()
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
