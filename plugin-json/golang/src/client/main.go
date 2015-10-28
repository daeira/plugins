package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

type Data map[string]interface{}
type Reply map[string]interface{}

func main() {

	// TODO: from discovery get by plugin name:
	//	 host:port
	//	 method -> name.method

	data := make(Data)
	reply, _ := call("Greeting.Process", "http://localhost:8881", data)
	fmt.Println("greeting", reply)
	for k, v := range reply {
		data[k] = v
	}
	reply, _ = call("Process", "http://localhost:8882", data)
	fmt.Println("reverse", reply)
	for k, v := range reply {
		data[k] = v
	}
	reply, _ = call("Process", "http://localhost:8883", data)
	fmt.Println("shuffle", reply)
	for k, v := range reply {
		data[k] = v
	}
	fmt.Println("data", data)

}

func call(name, url string, args map[string]interface{}) (map[string]interface{}, error) {
	msg, err := json2.EncodeClientRequest(name, []interface{}{args})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	//dreq, _ := httputil.DumpRequest(req, true)
	//log.Println(string(dreq))
	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//dresp, _ := httputil.DumpResponse(resp, true)
	//log.Println(string(dresp))
	reply := make(map[string]interface{})
	err = json2.DecodeClientResponse(resp.Body, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
