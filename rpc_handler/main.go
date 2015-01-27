package main

import (
    "fmt"
    "net"
)

func Handler(conn net.Conn) {
    fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

    buf := make([]byte, 1024)
    for {
        lenght, err := conn.Read(buf)
        if err != nil {
            conn.Close()
            break
        }
        if lenght > 0 {
            buf[lenght] = 0
        }
        fmt.Println("Rec[", conn.RemoteAddr().String(), "] Say :", string(buf[0:lenght]))

    }

}

func main() {
    if listener, err := net.Listen("tcp", ":3333"); err == nil {
        defer listener.Close()
        for {
            conn, err := listener.Accept()
            if err == nil {
                fmt.Println("Receive a Message")
                go Handler(conn)
            } else {
                fmt.Printf("Receive a err %s", err.Error())
            }
        }
    }
}
