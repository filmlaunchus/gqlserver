package gql

import (
  "github.com/graphql-go/graphql"
)

var (
  queryType
)

queryType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Query",
  Fields: graphql.Fields{
    "user": &userQuery,
  },
})
