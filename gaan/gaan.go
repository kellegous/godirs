package main

import (
  "github.com/kellegous/godirs"
  "os"
)

func main() {
  s, err := godirs.Run("/usr/local/go", "go")
  if err != nil {
    panic(err)
  }

  os.Exit(s)
}
