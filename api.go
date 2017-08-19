package main

import (
  "flag"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
)

var (
  addr = flag.String("addr", ":8000", "http service address")
  data map[string]string //Que significa?
)


func show(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  k:= p.ByName("key")
  if k == ""{
    fmt.Fprintf(w, "Read list: %v", data)
    return
  }
  fmt.Fprintf(w, "Read request: data[%s] = %s", k, data[k])
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  k:= p.ByName("key")
  v:= p.ByName("value")

  data[k] = v

  fmt.Fprintf(w, "Update: data[%s] = %s", k, data[k])
}

func main(){
  flag.Parse()
  data = map[string]string{}
  r:= httprouter.New()
  r.GET("/request/:key", show)
  r.GET("/list", show)
  r.PUT("/request/:key/:value", update)
  err := http.ListenAndServe(*addr, r)
  if err != nil{
    log.Fatal("ListenAndServe:", err)
  }
}
