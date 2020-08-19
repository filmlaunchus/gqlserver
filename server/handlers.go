/* handlers.go

*/
package server

import (
  "fmt"
  "net/http"
  "encoding/json"
  "context"
)

func (s *Server) GQLHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query().Get("query")
  enc   := json.NewEncoder(w)
  gqls.Run(guery, enc, s.objectStores)
}
