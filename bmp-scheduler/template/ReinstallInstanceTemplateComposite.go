package template

import (
	log "git.jd.com/cps-golang/log"
)

type ReinstallInstanceTemplateComposite struct{}

type ReinstallInstanceTemplater interface {
	accept(string, string) bool
	initCommand(*log.Logger, string, string, string, bool, bool, string, string)
}

var reinstallTemplates []ReinstallInstanceTemplater

func init() {
	reinstallTemplates = []ReinstallInstanceTemplater{}
	reinstallTemplates = append(reinstallTemplates, BasicReinstallInstanceTemplate{})
	// reinstallTemplates = append(reinstallTemplates, SDNBondReinstallInstanceTemplate{})
	// reinstallTemplates = append(reinstallTemplates, SDNDualReinstallInstanceTemplate{})
	// reinstallTemplates = append(reinstallTemplates, RetailBondReinstallInstanceTemplate{})
	// reinstallTemplates = append(reinstallTemplates, RetailDualReinstallInstanceTemplate{})
}

func (c ReinstallInstanceTemplateComposite) InitCommand(logger *log.Logger, network_type, interface_mode, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {
	for _, reinstallTemplate := range reinstallTemplates {
		if true { //reinstallTemplate.accept(network_type, interface_mode)
			reinstallTemplate.initCommand(logger, request_id, instance_id, sn, keep_data, make_partitions, image_format, task)
			break
		}
	}
}
