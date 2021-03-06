/* gql.go

package entry point

*/
package gql

import (
  "fmt"
  "encoding/json"

  "github.com/graphql-go/graphql"
)

type CRUDStore interface {
  Create(params map[string]interface{}) (interface{}, error)
  Read(id string) (interface{}, error)
  Update(id string, params map[string]interface{}) (interface{}, error)
  Delete(id string) (interface{}, error)
}

var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query:    queryType,
  Mutation: mutationType,
})

type GQL struct {
  schema graphql.Schema
}

func NewGQL() *GQL {
  return &GQL{schema: MainSchema}
}

func (gqls *GQL) Run(query string, enc json.Encoder, root map[string]interface{}) {
  result := graphql.Do(graphql.Params{
    RequestString: query,
    RootObject:    root,
    Schema:        gqls.schema,
  })

  if len(result.Errors) > 0 {
    fmt.Println("errors: %v", result.Errors)
  }
  enc.Encode(result)
}
