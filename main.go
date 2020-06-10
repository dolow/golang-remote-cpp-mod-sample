package main

import (
  "flag"
  "fmt"

  encryptor "github.com/dolow/golang-cpp-mod-sample"
)

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) == 0 {
    fmt.Printf("this module requires one argument\n")
    return
  }
  pt := args[0]

  var ct string

  encryptor.Encrypt(&ct, pt)

  ctbytes := []byte(ct)
  var ss string
  for c, _ := range ctbytes {
    ss += fmt.Sprintf("%0xd ", c)
  }
  fmt.Printf("ct: %s\n", ss)


  dt := encryptor.Decrypt(ct)
  fmt.Printf("dt: %s\n", dt)
}
