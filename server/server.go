/* server.go

*/
package server

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"

  ds  "github.com/filmlaunchus/gqlserver/datastores"
  gql "github.com/filmlaunchus/gqlserver/gql"
)

type Server struct {
    config        *Config
    objectStores  map[string]interface{}
    GQLEntry      *gql.GQL
}

// init server
func NewServer(configPath string) *Server {
  conf  := NewConfig(configPath)
  store := ds.InitStores(conf.mode)
  gqle  := gql.NewGQL()
  return &Server{
    config:        conf,
    objectStores:  store,
    GQLEntry:      gqle,
  }
}

// run server
func (s *Server) Run() {
  router := mux.NewRouter()

  router.HandleFunc("/gql", s.GQLHandler)
  http.Handle("/", router)

  fmt.Println("Server is running on port %s", s.config.port)
  http.ListenAndServe(":"+s.config.port, nil)
}
