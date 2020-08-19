/* gql-service.go

main entrypoint for server process
*/

package main

import (
  "fmt"

  "github.com/filmlaunchus/gqlserver/server"
)

func main() {
  s := server.NewServer(nil)
  s.Run()
}
