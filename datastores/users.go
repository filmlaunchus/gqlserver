/* datastores/users.go

*/
package datastores

import (
  // "fmt"
  "errors"

  "github.com/rs/xid"

  "github.com/filmlaunchus/gqlserver/utils"
)

// implements CRUDStore interface
type MockUserStore struct {
  users []*utils.UserObject
}

func NewMockUserStore() *MockUserStore {
  userSt := make([]*utils.UserObject, 0)
  return &MockUserStore{users: userSt}
}

func (mus *MockUserStore) Create(params map[string]interface{}) (interface{}, error) {
  uid   := xid.New()
  usern := params["username"].(string)
  email := params["email"].(string)

  uo := &utils.UserObject{uid.String(), usern, email}
  mus.users = append(mus.users, uo)
  return uo, nil
}

func (mus *MockUserStore) Read(id string) (interface{}, error) {
  for _, u := range mus.users {
    if id == u.Id {
      return u, nil
    }
  }
  return nil, errors.New("id not found")
}

func (mus *MockUserStore) Update(id string, params map[string]interface{}) (interface{}, error) {
  awsid, aok := params["username"].(string)
  email, eok := params["email"].(string)

  for _, u := range mus.users {
    if id == u.Id {
      if aok {
        u.Username = awsid
      }
      if eok {
        u.Email = email
      }
      return u, nil
    }
  }
  return nil, errors.New("id not found")
}

func (mus *MockUserStore) Delete(id string) (interface{}, error) {
  loc := -1
  uo  := &utils.UserObject{}

  for i, u := range mus.users {
    if id == u.Id {
      loc = i
      uo = u
      break
    }
  }
  if loc == -1 {
    return nil, errors.New("id not found")
  }
  mus.users[loc] = mus.users[len(mus.users)-1]
  mus.users = mus.users[:len(mus.users)-1]
  return uo, nil
}
