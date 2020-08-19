/* utils.go

*/
package utils

import (
  // "fmt"
)

type CRUDStore interface {
  Create(map[string]interface{}) (interface{}, error)
  Read(id string) (interface{}, error)
  Update(id string, map[string]interface{}) (interface{}, error)
  Delete(id string) (interface{}, error)
}

func adder(a, b int) int {
  return a+b
}
