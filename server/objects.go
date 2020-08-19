/* objects.go

*/
package server

import (
  "encoding/json"
  "time"
)

type UserObject struct {
  Id        string `json:"id"`
  Username  string `json:"username"`
  Email     string `json:"email"`
}
