/* handlers.go

*/
package server

import (
  "fmt"
  "net/http"
  "encoding/json"
  // "context"
  // ds "github.com/filmlaunchus/gqlserver/datastores"
)

func (s *Server) GQLHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("query")
  enc   := json.NewEncoder(w)
  fmt.Println("Running query: ", query)
  s.GQLEntry.Run(query, *enc, s.objectStores)
}
