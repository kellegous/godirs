package main

import (
  "github.com/kellegous/gotools"
  "os"
)

func main() {
  s, err := gotools.run("/usr/local/go", "godoc")
  if err != nil {
    panic(err)
  }

  os.Exit(s)
}
