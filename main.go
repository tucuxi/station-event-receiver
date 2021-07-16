package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"
)

var (
    delay int
    name string
    path string
    port int
)

func init() {
    flag.IntVar(&delay, "delay", 0, "Response delay in milliseconds")
    flag.StringVar(&name, "name", "stationid", "Station id parameter name")
    flag.StringVar(&path, "path", "/", "Request path")
    flag.IntVar(&port, "port", 8080, "Port number")
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
    stationId := req.FormValue(name)
    log.Println(req.Method, "Arrived at station", stationId)
    time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {
    flag.Parse()
    http.HandleFunc(path, handleRequest)
    address := fmt.Sprintf(":%d", port)
    if err := http.ListenAndServe(address, nil); err != nil {
        log.Fatal(err)
    }
}
