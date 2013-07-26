package main

import (
    "fmt"
    "log"
    "strconv"
    "encoding/json"
    "net/http"
    /* "code.google.com/p/go.net/websocket" */
)


type Page struct {
    WsUrl string `json:"webSocketDebuggerUrl"`
}


func connectToChrome(host string, port int, tab int) {

    address := host + ":" + strconv.Itoa(port) + "/json"
    resp, err := http.Get(address)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var pages []Page
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&pages)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(pages)
}

func main() {

    connectToChrome("http://localhost", 9222, 0)
}
