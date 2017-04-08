
# tool

Tool to compare size of json vs protobuf vs xml messages.

see nilsmagnus.github.io/ for article :

https://nilsmagnus.github.io/post/proto-json-sizes/

# usage

    go run main.go -show json -entries 200
    go run main.go -show jsonlen -entries 200
    go run main.go -show xmllen -entries 200
    go run main.go -show proto -entries 200
    go run main.go -show protolen -entries 200