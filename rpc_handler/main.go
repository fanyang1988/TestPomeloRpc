package main

import (
    "fmt"
    zmq "github.com/alecthomas/gozmq"
    sj "github.com/bitly/go-simplejson"
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
            continue
        }
        fmt.Printf("data  %s\n", parts)
        fmt.Printf("data 0  %s\n", parts[0])
        fmt.Printf("data 1  %s\n", parts[1])
        json, js_err := sj.NewJson(parts[1])

        if js_err != nil {
            fmt.Printf("js_err %s\n", js_err.Error())
            continue
        }
        /*
           {
             "id": 0,
             "msg": {
               "namespace": "user",
               "serverType": "test",
               "service": "service",
               "method": "echo",
               "args": [
                 "hello"
               ]
             }
           }
        */
        id, _ := json.Get("id").Int64()
        service_str, _ := json.Get("msg").Get("service").String()
        method_str, _ := json.Get("msg").Get("method").String()
        args, _ := json.Get("msg").Get("args").StringArray()

        fmt.Printf("data id  %d\n", id)
        fmt.Printf("data service  %s\n", service_str)
        fmt.Printf("data method  %s\n", method_str)
        fmt.Printf("data args  %s\n", args)

        resp_format := "{\"id\":%d,\"resp\":[%s, \"%s\"]}"
        null_str := "null"

        resp := fmt.Sprintf(resp_format, id, null_str, args[0])

        router.SendMultipart([][]byte{parts[0], []byte(resp)}, 0)
    }
}
