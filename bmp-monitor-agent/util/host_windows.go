// +build windows

package util

import (
	"context"
	"errors"
	"time"

	"github.com/StackExchange/wmi"
)

var (
	Timeout    = 30 * time.Second
	ErrTimeout = errors.New("command timed out")
)

type Win32_BIOS struct {
	Manufacturer string
	SerialNumber string
}

func SN() (HostInfo, error) {
	var dst []Win32_BIOS

	ret := HostInfo{}
	err := WMIQuery("SELECT * FROM Win32_BIOS", &dst)
	if err != nil {
		return ret, err
	}

	for _, d := range dst {
		if len(d.SerialNumber) == 0 {
			continue
		}

		ret.Manufacturer = d.Manufacturer
		ret.SerialNumber = d.SerialNumber
	}
	if ret.SerialNumber == "" {
		ret.SerialNumber = "empty_sn_mocked" //有些虚拟机没有sn时mock，测试方便。TODO 上线前放开
	}
	return ret, nil
}

func WMIQuery(query string, dst interface{}, connectServerArgs ...interface{}) error {
	return WMIQueryWithContext(context.Background(), query, dst, connectServerArgs...)
}

// WMIQueryWithContext - wraps wmi.Query with a timed-out context to avoid hanging
func WMIQueryWithContext(ctx context.Context, query string, dst interface{}, connectServerArgs ...interface{}) error {
	if _, ok := ctx.Deadline(); !ok {
		ctxTimeout, cancel := context.WithTimeout(ctx, Timeout)
		defer cancel()
		ctx = ctxTimeout
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.Query(query, dst, connectServerArgs...)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
