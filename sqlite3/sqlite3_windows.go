/*=================================
* Copyright(c)2015-2016 gostores
* All rights reserved
* Inspired by mattn/go-sqlite3
*=================================*/

package sqlite3

/*
#cgo CFLAGS: -I. -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo windows,386 CFLAGS: -D_USE_32BIT_TIME_T
#cgo LDFLAGS: -lmingwex -lmingw32
*/
import "C"
