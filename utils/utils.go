/* utils.go

*/
package utils

import (
  // "fmt"
)

type CRUDStore interface {
  Create(p map[string]interface{}) (interface{}, error)
  Read(id string) (interface{}, error)
  Update(id string, p map[string]interface{}) (interface{}, error)
  Delete(id string) (interface{}, error)
}

func adder(a int, b int) int {
  return a+b
}
