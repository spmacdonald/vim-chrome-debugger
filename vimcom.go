package main


import (
    "log"
    "net"
)


func handleConnection(c net.Conn) {
    buf := make([]byte, 1024)
    n, err := c.Read(buf)
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Print(string(buf[:n]))
}


func main() {
    ln, err := net.Listen("tcp", ":3219")
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
            continue
        }
        go handleConnection(conn)
    }
}
