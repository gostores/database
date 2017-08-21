/*=================================
* Copyright(c)2015-2016 Yiqilai
* All rights reserved
* Inspired by go-sql-driver/mysql
*=================================*/

package mysql

import (
	"bytes"
	"log"
	"testing"
)

func TestErrorsSetLogger(t *testing.T) {
	previous := errLog
	defer func() {
		errLog = previous
	}()

	// set up logger
	const expected = "prefix: test\n"
	buffer := bytes.NewBuffer(make([]byte, 0, 64))
	logger := log.New(buffer, "prefix: ", 0)

	// print
	SetLogger(logger)
	errLog.Print("test")

	// check result
	if actual := buffer.String(); actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestErrorsStrictIgnoreNotes(t *testing.T) {
	runTests(t, dsn+"&sql_notes=false", func(dbt *DBTest) {
		dbt.mustExec("DROP TABLE IF EXISTS does_not_exist")
	})
}
