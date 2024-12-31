// +build darwin

package util

import (
	"os/exec"
	"strings"
)

func SN() (HostInfo, error) {
	ret := HostInfo{}

	ioreg, err := exec.LookPath("/usr/sbin/ioreg")
	if err != nil {
		return ret, err
	}
	cmd := exec.Command(ioreg, "-l")

	out, err := cmd.Output()
	if err != nil {
		return ret, err
	}
	for _, line := range strings.Split(string(out), "\n") {
		if len(ret.SerialNumber) == 0 && strings.Contains(line, "\"IOPlatformSerialNumber\"") {
			s := strings.Split(line, "=")
			ret.SerialNumber = strings.ReplaceAll(s[len(s)-1], "\"", "")
			println("SerialNumber:" + ret.SerialNumber)
			continue
		}

		if len(ret.Manufacturer) == 0 && strings.Contains(line, "\"Manufacturer\"") {
			s := strings.Split(line, "=")
			ret.Manufacturer = strings.ReplaceAll(s[len(s)-1], "\"", "")
			println("Manufacturer:" + ret.Manufacturer)
			continue
		}
	}
	return ret, nil
}
