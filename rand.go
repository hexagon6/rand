package main

import (
  "time"
  "fmt"
  "math/rand"
)

func main() {
  t := time.Now().UnixNano()
  r := rand.New(rand.NewSource(t))
  fmt.Println(r.Float64())
}

