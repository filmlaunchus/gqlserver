/* datastores.go

*/
package datastores

import (
  // "fmt"

  // "github.com/filmlaunchus/gqlserver/utils"
)

type Datastore struct {
  objectStores map[string]interface{}
}

func NewDatastore(mode string) *Datastore {
  oS := make(map[string]interface{})
  if mode == "mock" {
    oS["users"] = NewMockUserStore()
  }
  return &Datastore{objectStores: oS}
}
