package template

import (
	log "git.jd.com/cps-golang/log"
)

type DeleteInstanceTemplateComposite struct{}

type DeleteInstanceTemplater interface {
	accept(string) bool
	initCommand(*log.Logger, string, string, string, string)
}

var deletetemplates []DeleteInstanceTemplater

func init() {
	deletetemplates = []DeleteInstanceTemplater{}
	deletetemplates = append(deletetemplates, BasicDeleteInstanceTemplate{})
	deletetemplates = append(deletetemplates, SDNDeleteInstanceTemplate{})
	deletetemplates = append(deletetemplates, RetailDeleteInstanceTemplate{})

}

func (c DeleteInstanceTemplateComposite) InitCommand(logger *log.Logger, network_type, request_id, instance_id, sn string, task string) {
	for _, dtemplate := range deletetemplates {
		//if dtemplate.accept(network_type) {
		dtemplate.initCommand(logger, request_id, instance_id, sn, task)
		break
		//}
	}
}
