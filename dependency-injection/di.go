package main

import (
  "fmt"
  "bytes"
)

func Greet( writer *bytes.Buffer, name string) string {
  fmt.Fprintf(writer, "Hello, %s", name)
  return name
}

