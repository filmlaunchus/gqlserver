/* datastores.go

*/
package datastores

import (
  // "fmt"

  // "github.com/filmlaunchus/gqlserver/utils"
)

func InitStores(mode string) map[string]interface{} {
  oS := make(map[string]interface{})
  if mode == "mock" {
    oS["users"] = NewMockUserStore()
  }
  return oS
}
