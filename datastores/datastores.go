/* datastores.go

*/
package datastores

import (
  "fmt"
)

type CRUDStore interface {
  Create(map[string]interface{}) (interface{}, error)
  Read(id string) (interface{}, error)
  Update(id string, map[string]interface{}) (interface{}, error)
  Delete(id string) (interface{}, error)
}

type Datastore struct {
  objectStores map[string]CRUDStore
}

func NewDatastore(mode string) *Datastore {
  oS := make(map[string]CRUDStore)
  if mode == "mock" {
    oS["users"] = NewMockUserStore()
  }
  return &Datastore{objectStores: oS}
}
