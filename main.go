package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nilsmagnus/protobuf-json-xml-size-comparison/sample"
)

func main() {
	fmt.Println("| json | gzipped json | proto | gzipped proto | proto size(%) of json | gzipped proto size(%) of gzipped json |")
	for _, dataSize := range []int{0, 1, 2, 10, 20, 200, 2000, 20000} {
		protoStruct := createTestDatata(dataSize)
		jsonl, gzJsonlen, protol, gzProto := jsonProtoLengts(protoStruct)
		fmt.Printf("| %d |  %d | %d | %d | %d | %f | %f | \n", dataSize, gzJsonlen, protol, gzProto, gzProto, float32(gzProto)/float32(gzJsonlen), float32(protol)/float32(jsonl))
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

func jsonProtoLengts(protoSome *sample.Test) (jsonLen, gzipJSONLen, protoLen, gzpProto int) {
	data, _ := proto.Marshal(protoSome)
	protoLen = len(data)
	jsonified, _ := json.Marshal(protoSome)
	jsonLen = len(jsonified)
	gzipJSONLen = gzipLen(jsonified)
	gzpProto = gzipLen(data)
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
