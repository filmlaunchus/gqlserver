/* objects.go

*/
package utils

import (
  // "encoding/json"
)

type UserObject struct {
  Id        string `json:"id"`
  Username  string `json:"username"`
  Email     string `json:"email"`
}
