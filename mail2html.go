package main

import (
  "github.com/atotto/clipboard"
  "io/ioutil"
  "mime/quotedprintable"
  "net/mail"
  "strings"
)

func main() {
  source, _ := clipboard.ReadAll()
  reader := strings.NewReader(source)
  email, _ := mail.ReadMessage(reader)
  body, _ := ioutil.ReadAll(email.Body)

  html, _ := ioutil.ReadAll(quotedprintable.NewReader(strings.NewReader(string(body))))

  _ = clipboard.WriteAll(string(html))
}
