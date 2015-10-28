package rpcsvc

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

func Serve(name, addr string, rcvr interface{}) {
	mux := http.NewServeMux()
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(rcvr, name)
	mux.Handle("/", s)

	// TODO: register Plugin with host:port method
	log.Println("serving on", addr)
	graceful.Run(addr, 10*time.Second, mux)
	// TODO: deregister Plugin with host:port method
}
