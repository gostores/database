/*=================================
* Copyright(c)2015-2016 Yiqilai
* All rights reserved
* Inspired by go-sql-driver/mysql
*=================================*/

package mysql

import (
	"database/sql/driver"
	"io"
)

type mysqlField struct {
	tableName string
	name      string
	flags     fieldFlag
	fieldType byte
	decimals  byte
}

type mysqlRows struct {
	mc      *mysqlConn
	columns []mysqlField
}

type binaryRows struct {
	mysqlRows
}

type textRows struct {
	mysqlRows
}

type emptyRows struct{}

func (rows *mysqlRows) Columns() []string {
	columns := make([]string, len(rows.columns))
	if rows.mc.cfg.columnsWithAlias {
		for i := range columns {
			columns[i] = rows.columns[i].tableName + "." + rows.columns[i].name
		}
	} else {
		for i := range columns {
			columns[i] = rows.columns[i].name
		}
	}
	return columns
}

func (rows *mysqlRows) Close() error {
	mc := rows.mc
	if mc == nil {
		return nil
	}
	if mc.netConn == nil {
		return ErrInvalidConn
	}

	// Remove unread packets from stream
	err := mc.readUntilEOF()
	rows.mc = nil
	return err
}

func (rows *binaryRows) Next(dest []driver.Value) error {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows *textRows) Next(dest []driver.Value) error {
	if mc := rows.mc; mc != nil {
		if mc.netConn == nil {
			return ErrInvalidConn
		}

		// Fetch next row from stream
		return rows.readRow(dest)
	}
	return io.EOF
}

func (rows emptyRows) Columns() []string {
	return nil
}

func (rows emptyRows) Close() error {
	return nil
}

func (rows emptyRows) Next(dest []driver.Value) error {
	return io.EOF
}
