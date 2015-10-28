package main

import (
	"fmt"
	"net"
	"net/http"
	"rpcsvc"
)

type Greeting struct {
}

func (a *Greeting) Process(r *http.Request, data *map[string]interface{}, reply *map[string]interface{}) error {
	fmt.Println("Greeting.Process args:", data)
	*reply = make(map[string]interface{})
	(*reply)["greeting"] = "Nice to meet you"
	fmt.Println("Greeting.Process reply:", reply)
	return nil
}

func main() {

	// TODO: with service discovery:
	//port, err := freePort()
	//if err != nil {
	//log.Fatal(err)
	//}

	rpcsvc.Serve("Greeting", "127.0.0.1:8881", new(Greeting))
}

func freePort() (string, error) {
	ln, _ := net.Listen("tcp", "")
	defer ln.Close()
	_, port, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		return "", err
	}
	return port, nil
}
