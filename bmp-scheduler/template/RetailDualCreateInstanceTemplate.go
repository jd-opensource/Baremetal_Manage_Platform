package template

import (
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InterfaceMode"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type RetailDualCreateInstanceTemplate struct {
	AbstractRetailCreateInstanceTemplate
}

func (b RetailDualCreateInstanceTemplate) accept(network_type, interface_mode string) bool {
	return NetworkType.RETAIL == network_type && interface_mode == InterfaceMode.DUAL
}

func (b RetailDualCreateInstanceTemplate) createSetNetworkCommand(tx *gorm.DB, request_id, instance_id, sn string, parent_command_id int64, task string) (int64, error) {
	return CreateCommand(tx, CommandAction.SetNetwork.Name, CommandType.AGENT, request_id, instance_id, sn, parent_command_id, task)
}

func (b RetailDualCreateInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, make_partitions bool, image_format string, task string) {
	logger.Info("RetailDualCreateInstanceTemplate start", request_id, instance_id, sn)
	b.AbstractRetailCreateInstanceTemplate.createSetNetworkCommand = b.createSetNetworkCommand
	b.AbstractRetailCreateInstanceTemplate.initCommand(logger, request_id, instance_id, sn, make_partitions, image_format, task)
}
