/* users.go

*/

package gql

import (
  // "fmt"

  "github.com/graphql-go/graphql"

  "github.com/filmlaunchus/gqlserver/utils"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
  Name: "User",
  Fields: graphql.Fields{
    "id": &graphql.Field{
      Type: graphql.NewNonNull(graphql.String),
      Description: "User GUID",
      Resolve: func (p graphql.ResolveParams) (interface{}, error) {
        if user, ok := p.Source.(utils.UserObject); ok {
          return user.Id, nil
        }
        return nil, nil
      },
    },
    "username": &graphql.Field{
      Type: graphql.String,
      Description: "User's username",
      Resolve: func (p graphql.ResolveParams) (interface{}, error) {
        if user, ok := p.Source.(utils.UserObject); ok {
          return user.Username, nil
        }
        return nil, nil
      },
    },
    "email": &graphql.Field{
      Type: graphql.String,
      Description: "User Email",
      Resolve: func (p graphql.ResolveParams) (interface{}, error) {
        if user, ok := p.Source.(utils.UserObject); ok {
          return user.Email, nil
        }
        return nil, nil
      },
    },
  },
})

var userQuery = graphql.Field{
  Type: userType,
  Args: graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
      Description: "id of user that you wanna get",
      Type: graphql.NewNonNull(graphql.String),
    },
  },
  Resolve: func (p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]CRUDStore)
    user, err := objs["users"].Read(p.Args["id"])
    return user, err
  },
}

var createUserMut = graphql.Field{
  Type: userType,
  Description: "Create new user",
  Args: graphql.FieldConfigArgument{
    "username": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.String),
    },
    "email": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.String),
    },
  },
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].Create(p.Args)
    return user, err
  },
}

var updateUserMut = graphql.Field{
  Type: userType,
  Description: "Update user by id",
  Args: graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.String),
    },
    "email": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.String),
    },
  },
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].Update(p.Args)
    return user, err
  },
}

var deleteUserMut = graphql.Field{
  Type: userType,
  Description: "Delete product by id",
  Args: graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.Int),
    },
  },
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].Delete(p.Args)
    return user, err
  },
}
