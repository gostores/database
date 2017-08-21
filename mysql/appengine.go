// +build appengine

package mysql

import (
	"appengine/cloudsql"
)

func init() {
	RegisterDial("cloudsql", cloudsql.Dial)
}
