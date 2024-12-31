// +build windows

package util

import "errors"

var ErrNotImplemented = errors.New("Not implemented")

func Daemon() error {
	return ErrNotImplemented
}
