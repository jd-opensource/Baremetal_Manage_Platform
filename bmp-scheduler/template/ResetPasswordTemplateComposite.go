package template

import (
	log "git.jd.com/cps-golang/log"
)

type ResetPasswordTemplateComposite struct{}

type ResetPasswordTemplater interface {
	accept(string) bool
	initCommand(*log.Logger, string, string, string, string)
}

var resetPasswordtemplates []ResetPasswordTemplater

func init() {
	resetPasswordtemplates = []ResetPasswordTemplater{}
	resetPasswordtemplates = append(resetPasswordtemplates, BasicResetPasswordTemplate{})
	// resetPasswordtemplates = append(resetPasswordtemplates, SDNResetPasswordTemplate{})
	// resetPasswordtemplates = append(resetPasswordtemplates, RetailResetPasswordTemplate{})

}

func (c ResetPasswordTemplateComposite) InitCommand(logger *log.Logger, network_type, request_id, instance_id, sn string, task string) {
	for _, rtemplate := range resetPasswordtemplates {
		if true { //rtemplate.accept(network_type)
			rtemplate.initCommand(logger, request_id, instance_id, sn, task)
			break
		}
	}
}
