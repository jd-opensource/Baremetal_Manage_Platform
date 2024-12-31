package template

import (
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InterfaceMode"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type SDNBondReinstallInstanceTemplate struct {
	AbstractSDNReinstallInstanceTemplate
}

func (b SDNBondReinstallInstanceTemplate) accept(network_type, interface_mode string) bool {
	return NetworkType.VPC == network_type && interface_mode != InterfaceMode.DUAL

}

func (b SDNBondReinstallInstanceTemplate) createSetNetworkCommand(tx *gorm.DB, request_id, instance_id, sn string, parent_command_id int64, task string) (int64, error) {

	return CreateCommand(tx, CommandAction.SetBond.Name, CommandType.AGENT,
		request_id, instance_id, sn, parent_command_id, task)
}

func (b SDNBondReinstallInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {
	b.AbstractSDNReinstallInstanceTemplate.createSetNetworkCommand = b.createSetNetworkCommand
	b.AbstractSDNReinstallInstanceTemplate.initCommand(logger, request_id, instance_id, sn, keep_data, make_partitions, image_format, task)
	logger.Point("SDNBondReinstallInstanceTemplate_status", "finish")
}
