package main

import (
  "github.com/kellegous/gotools"
  "os"
)

func main() {
  s, err := gotools.Run("/usr/local/go", "godoc")
  if err != nil {
    panic(err)
  }

  os.Exit(s)
}
