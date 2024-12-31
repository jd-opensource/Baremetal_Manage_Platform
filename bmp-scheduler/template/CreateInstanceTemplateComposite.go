package template

import (
	"fmt"

	log "git.jd.com/cps-golang/log"
)

type CreateInstanceTemplateComposite struct{}

type CreateInstanceTemplater interface {
	accept(string, string) bool
	initCommand(*log.Logger, string, string, string, bool, string, string)
}

var createtemplates []CreateInstanceTemplater

func init() {
	createtemplates = []CreateInstanceTemplater{}
	createtemplates = append(createtemplates, BasicCreateInstanceTemplate{})
	//createtemplates = append(createtemplates, SDNBondCreateInstanceTemplate{})
	//createtemplates = append(createtemplates, SDNDualCreateInstanceTemplate{})
	//createtemplates = append(createtemplates, RetailBondCreateInstanceTemplate{})
	//createtemplates = append(createtemplates, RetailDualCreateInstanceTemplate{})
}

func (c CreateInstanceTemplateComposite) InitCommand(logger *log.Logger, network_type, interface_mode, request_id, instance_id, sn string, make_partitions bool, image_format string, task string) {
	fmt.Println(createtemplates, network_type, interface_mode)
	for _, ctemplate := range createtemplates {
		if true { //ctemplate.accept(network_type, interface_mode) {
			ctemplate.initCommand(logger, request_id, instance_id, sn, make_partitions, image_format, task)
			break
		}
	}
}
