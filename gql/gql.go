/* gql.go

package entry point

*/
package gql

import (
  "fmt"
  "encoding/json"

  "github.com/graphql-go/graphql"
)

var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query:    queryType,
  Mutation: mutationType,
})

type GQL struct {
  schema graphql.Schema
}

func NewGQL() *GQL {
  return &GQL{
    schema: MainSchema
  }
}

func (gqls *GQL) Do(guery string, enc Encoder, root map[string]interface{}) {
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
