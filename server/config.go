/* config.go

*/
package server

import (
  "fmt"
  "io/ioutil"

  "github.com/buger/jsonparser"
)

var (
  defaultPort = "8080"
  defaultMode = "mock"
)

type Config struct {
  port string
  mode string
}

func NewConfig(cP string) *Config {
  conf := &Config{
    port: defaultPort,
    mode: defaultMode
  }
  if cP != nil {
    data, err := ioutil.ReadFile(cP)
    if err != nil {
      fmt.Println("File reading error", err)
      return
    }
    if val, err := jsonparser.GetString(data, "config", "mode"); err != nil {
      conf.mode = val
    }
    if val, err := jsonparser.GetString(data, "config", "port"); err != nil {
      conf.port = val
    }
  }
  return conf
}
