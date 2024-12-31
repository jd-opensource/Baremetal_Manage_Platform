// +build linux

package util

import (
	"os/exec"
	"strings"
)

func SN() (HostInfo, error) {
	ret := HostInfo{}

	dmidcode, err := exec.LookPath("dmidecode")
	if err != nil {
		return ret, err
	}
	cmd := exec.Command(dmidcode, "-t", "1")

	out, err := cmd.Output()
	if err != nil {
		return ret, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		tmp := strings.SplitN(line, ":", 2)
		if strings.EqualFold(strings.TrimSpace(tmp[0]), "Serial Number") {
			ret.SerialNumber = strings.TrimSpace(tmp[1])
			continue
		}

		if strings.EqualFold(strings.TrimSpace(tmp[0]), "Manufacturer:") {
			ret.Manufacturer = strings.TrimSpace(tmp[1])
			continue
		}
	}

	return ret, nil
}
