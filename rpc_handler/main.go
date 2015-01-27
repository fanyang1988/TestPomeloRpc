package main

import (
    "fmt"
    zmq "github.com/alecthomas/gozmq"
)

func main() {
    context, _ := zmq.NewContext()
    defer context.Close()

    router, _ := context.NewSocket(zmq.ROUTER)
    defer router.Close()
    router.Bind("tcp://*:3333")

    for {
        parts, err := router.RecvMultipart(0)

        if err != nil {
            fmt.Printf("err %s\n", err.Error())
        }
        fmt.Printf("data  %s\n", parts)
        fmt.Printf("data 0  %s\n", parts[0])
        fmt.Printf("data 1  %s\n", parts[1])

        router.SendMultipart([][]byte{parts[0], []byte("{\"id\":0,\"resp\":[null, \"22222\"]}")}, 0)
    }
}
