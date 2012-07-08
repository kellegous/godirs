package main

import (
  "github.com/kellegous/godirs"
  "os"
)

func main() {
  root, err := godirs.FindGoRoot()
  if err != nil {
    panic(err)
  }

  s, err := godirs.Run(root, "godoc")
  if err != nil {
    panic(err)
  }

  os.Exit(s)
}
