/* gql-service.go

main entrypoint for server process
*/

package gqlserver

import (
  "fmt"

  "github.com/filmlaunchus/gql-service/server"
)

func main() {
  s := server.NewServer(nil)
  s.Run()
}
