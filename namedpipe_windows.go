//go:build windows
// +build windows

package mssql

import (
	"context"
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
	"net"
	"time"

	"github.com/BruceCatYu/go-mssqldb/msdsn"
)

func dialConnectionUsingNamedPipe(_ context.Context, _ *Connector, p msdsn.Config) (conn net.Conn, err error) {
	conn, err = npipe.DialTimeout(p.Host, 3*time.Second)

	if err != nil {
		f := "Unable to open named pipe connection with address '%v': %v"
		return nil, fmt.Errorf(f, p.Host, err.Error())
	}

	return conn, err
}
