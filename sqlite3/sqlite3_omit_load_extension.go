/*=================================
* Copyright(c)2015-2016 gostores
* All rights reserved
* Inspired by mattn/go-sqlite3
*=================================*/

package sqlite3

/*
#cgo CFLAGS: -DSQLITE_OMIT_LOAD_EXTENSION
*/
import "C"
import (
	"errors"
)

func (c *SQLiteConn) loadExtensions(extensions []string) error {
	return errors.New("Extensions have been disabled for static builds")
}
