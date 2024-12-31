package template

import (
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InterfaceMode"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type RetailDualReinstallInstanceTemplate struct {
	AbstractRetailReinstallInstanceTemplate
}

func (b RetailDualReinstallInstanceTemplate) accept(network_type, interface_mode string) bool {
	return NetworkType.RETAIL == network_type && interface_mode == InterfaceMode.DUAL
}

func (b RetailDualReinstallInstanceTemplate) createSetNetworkCommand(tx *gorm.DB, request_id, instance_id, sn string, parent_command_id int64, task string) (int64, error) {
	return CreateCommand(tx, CommandAction.SetNetwork.Name, CommandType.AGENT, request_id, instance_id, sn, parent_command_id, task)
}

func (b RetailDualReinstallInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {
	b.AbstractRetailReinstallInstanceTemplate.createSetNetworkCommand = b.createSetNetworkCommand
	b.AbstractRetailReinstallInstanceTemplate.initCommand(logger, request_id, instance_id, sn, keep_data, make_partitions, image_format, task)
}
