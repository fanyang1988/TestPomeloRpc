package main

import (
    "fmt"
    zmq "github.com/alecthomas/gozmq"
    //sj "github.com/bitly/go-simplejson"
    "time"
)

// add by b2
// add by b3
//
//
// add by b4
//

func main() {
    context, _ := zmq.NewContext()
    defer context.Close()

    router, _ := context.NewSocket(zmq.DEALER)
    defer router.Close()

    router.SetIdentity("test-server-1")
    router.Connect("tcp://localhost:3333")
    id := 1

    for {
        time.Sleep(1 * time.Second)
        data_format := `
{
  "id": %d,
  "msg": {
    "namespace": "user",
    "serverType": "test",
    "service": "service",
    "method": "echo",
    "args": [
      "%s"
    ]
  }
}
`
        msg := fmt.Sprintf("hello b1 b2 %d", id)
        data := fmt.Sprintf(data_format, id, msg)
        id++
        router.Send([]byte(data), 0)

        parts, err := router.RecvMultipart(0)

        if err != nil {
            fmt.Printf("err %s\n", err.Error())
            continue
        }
        fmt.Printf("data  %s\n", parts)
        fmt.Printf("data 0  %s\n", parts[0])
    }
}
