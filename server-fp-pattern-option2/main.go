package main

import (
  "log"
  "time"
)

func main() {
  svr := server.New(server.Config{"localhost", 1234, time.Second, 10})
  if err := svr.Start(); err != nil {
    log.Fatal(err)
  }
}
