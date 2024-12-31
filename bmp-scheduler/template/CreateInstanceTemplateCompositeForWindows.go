package template

import (
	"fmt"

	log "git.jd.com/cps-golang/log"
)

type CreateInstanceTemplateCompositeForWindows struct{}

type CreateInstanceTemplaterForWindows interface {
	accept(string, string) bool
	initCommand(*log.Logger, string, string, string, bool, string, string)
}

var createtemplatesForWindows []CreateInstanceTemplaterForWindows

func init() {
	createtemplatesForWindows = []CreateInstanceTemplaterForWindows{}
	createtemplatesForWindows = append(createtemplatesForWindows, BasicCreateInstanceTemplateForWindows{})
	//createtemplatesForWindows = append(createtemplatesForWindows, SDNBondCreateInstanceTemplate{})
	//createtemplatesForWindows = append(createtemplatesForWindows, SDNDualCreateInstanceTemplate{})
	//createtemplatesForWindows = append(createtemplatesForWindows, RetailBondCreateInstanceTemplate{})
	//createtemplatesForWindows = append(createtemplatesForWindows, RetailDualCreateInstanceTemplate{})
}

func (c CreateInstanceTemplateCompositeForWindows) InitCommand(logger *log.Logger, network_type, interface_mode, request_id, instance_id, sn string, make_partitions bool, image_format string, task string) {
	fmt.Println(createtemplatesForWindows, network_type, interface_mode)
	for _, ctemplate := range createtemplatesForWindows {
		if true { //ctemplate.accept(network_type, interface_mode) {
			ctemplate.initCommand(logger, request_id, instance_id, sn, make_partitions, image_format, task)
			break
		}
	}
}
