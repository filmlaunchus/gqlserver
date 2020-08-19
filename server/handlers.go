/* handlers.go

*/
package server

import (
  // "fmt"
  "net/http"
  "encoding/json"
  // "context"
  // ds "github.com/filmlaunchus/gqlserver/datastores"
)

func (s *Server) GQLHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("query")
  enc   := json.NewEncoder(w)
  root  := s.objectStores
  s.GQLEntry.Run(query, *enc, root)
}
