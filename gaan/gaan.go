package main

import "os"

func main() {
  s, err := run("/usr/local/go", "go")
  if err != nil {
    panic(err)
  }

  os.Exit(s)
}
