// +build appengine

/*=================================
* Copyright(c)2015-2016 gostores
* All rights reserved
* Inspired by go-sql-driver/mysql
*=================================*/

package mysql

import (
	"appengine/cloudsql"
)

func init() {
	RegisterDial("cloudsql", cloudsql.Dial)
}
