package dgraph

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/dgraph-io/gru/gruadmin/server"
)

var (
	Server = flag.String("dgraph", "http://127.0.0.1:8080", "Dgraph server address")
	// TODO - Remove this later.
	QueryEndpoint = strings.Join([]string{*Server, "query"}, "/")
	endpoint      = strings.Join([]string{*Server, "query"}, "/")
)

type MutationRes struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Uids    map[string]string `json:"uids"`
}

func SendMutation(m string) server.Response {
	res, err := http.Post(endpoint, "application/x-www-form-urlencoded", strings.NewReader(m))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var mr MutationRes
	json.NewDecoder(res.Body).Decode(&mr)
	fmt.Println(mr)
	if mr.Code == "ErrorOk" {
		return server.Response{
			Success: true,
		}
	}
	return server.Response{
		Success: false,
	}
}

func Query(q string) []byte {
	res, err := http.Post(endpoint, "application/x-www-form-urlencoded", strings.NewReader(q))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
