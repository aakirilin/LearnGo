package sse

import "github.com/alexandrevicenzi/go-sse"

func creteSSEServer() *sse.Server {
	s := sse.NewServer(nil)
	return s
}

var SseServer = creteSSEServer()
