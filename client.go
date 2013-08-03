package foo

import (
    /* "fmt" */
    "log"
    "strconv"
    "encoding/json"
    "net/http"
    "code.google.com/p/go.net/websocket"
)


type Page struct {
    WsUrl string `json:"webSocketDebuggerUrl"`
}


type Request struct {
    Id int `json:"id"`
    Method string `json:"method"`
    Params map[string]interface{} `json:"params"`
}


type Response struct {
    Id int `json:"id"`
    Error interface{} `json:"Error"`
}


func connectToChrome(host string, port int, tab int) *websocket.Conn {

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

    url := pages[tab].WsUrl
    ws, err := websocket.Dial(url, "", "http://localhost/")
    if err != nil {
        log.Fatal(err)
    }

    return ws
}


func listenWs(ws *websocket.Conn) {
    for {
        var x interface{}
        err := websocket.JSON.Receive(ws, &x)
        if err != nil {
            log.Fatal(err)
        }

        log.Print(x)
    }

    ws.Close()
}


func main() {

    ws := connectToChrome("http://localhost", 9222, 0)
    go listenWs(ws)

    req := Request{Id: 1, Method: "Page.enable"}
    /* req := Request{Id: 1, Method: "Page.reload"} */
    websocket.JSON.Send(ws, req)

    var resp Response
    websocket.JSON.Receive(ws, &resp)

    log.Print(resp)

    select {}
}
