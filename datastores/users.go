/* datastores/users.go

*/
package datastores

import (
  "fmt"
  "errors"

  "github.com/rs/xid"

  "../utils"
)

// implements CRUDStore interface
type MockUserStore struct {
  users []*utils.UserObject
}

func NewMockUserStore() *MockUserStore {
  usersSt := make([]*utils.UserObject)
  return &MockUserStore{users: userSt}
}

func (mus *MockUserStore) Create(params map[string]interface{}) (interface{}, error) {
  uid   := xid.New()
  usern := params["username"].(string)
  email := params["email"].(string)

  uo := &utils.UserObject{uid, usern, email}
  mus.userSt = append(mus.userSt, uo)
  return uo, nil
}

func (mus *MockUserStore) Read(id string) (interface{}, error) {
  for _, u := range mus.userSt {
    if uid == u.Id {
      return u, nil
    }
  }
  return nil, errors.New("%s not found", uid)
}

func (mus *MockUserStore) Update(id string, params map[string]interface{}) (interface{}, err) {
  awsid, aok := params["username"].(string)
  email, eok := params["email"].(string)

  for _, u := range mus.userSt {
    if uid == u.Id {
      if aok {
        u.Username = awsid
      }
      if eok {
        u.Email = email
      }
      return u, nil
    }
  }
  return nil, errors.New("%s not found", uid)
}

func (mus *MockUserStore) Delete(id string) (interface{}, err) {
  loc := -1
  uo  := &utils.UserObject{}

  for i, u := range mus.userSt {
    if id == u.Id {
      loc = i
      uo = u
      break
    }
  }
  if loc == -1 {
    return nil, errors.New("%s not found", id)
  }
  mus.UserSt[loc] = mus.UserSt[len(mus.UserSt)-1]
  mus.UserSt = mus.UserSt[:len(mus.UserSt)-1]
  return uo, nil
}
