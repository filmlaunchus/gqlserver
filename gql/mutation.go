/* mutation.go

*/
package gql

import (
  "github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &createUserMut,
		"updateUser": &updateUserMut,
		"deleteUser": &deleteUserMut,
	},
})
