/* users.go

*/

package gql

import (
  "fmt"

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
        if user, ok := p.Source.(*utils.UserObject); ok {
          return user.Id, nil
        }
        fmt.Println("not a UO??")
        return nil, nil
      },
    },
    "username": &graphql.Field{
      Type: graphql.String,
      Description: "User's username",
      Resolve: func (p graphql.ResolveParams) (interface{}, error) {
        if user, ok := p.Source.(*utils.UserObject); ok {
          return user.Username, nil
        }
        return nil, nil
      },
    },
    "email": &graphql.Field{
      Type: graphql.String,
      Description: "User Email",
      Resolve: func (p graphql.ResolveParams) (interface{}, error) {
        if user, ok := p.Source.(*utils.UserObject); ok {
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
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].(CRUDStore).Read(p.Args["id"].(string))
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
    user, err := objs["users"].(CRUDStore).Create(p.Args)
    fmt.Println("created new user object", user)
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
      Type: graphql.String,
    },
    "username": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].(CRUDStore).Update(p.Args["id"].(string), p.Args)
    return user, err
  },
}

var deleteUserMut = graphql.Field{
  Type: userType,
  Description: "Delete product by id",
  Args: graphql.FieldConfigArgument{
    "id": &graphql.ArgumentConfig{
      Type: graphql.NewNonNull(graphql.String),
    },
  },
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    objs := p.Info.RootValue.(map[string]interface{})
    user, err := objs["users"].(CRUDStore).Delete(p.Args["id"].(string))
    return user, err
  },
}
