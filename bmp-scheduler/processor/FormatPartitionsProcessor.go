package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instancePartitionDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type FormatPartitionsProcessor struct {
	BaseProcessor
}

func NewFormatPartitionsProcessor() FormatPartitionsProcessor {
	b := FormatPartitionsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b FormatPartitionsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b FormatPartitionsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b FormatPartitionsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("FormatPartitionsProcessor doProcess......", command.Sn)
	instancePartitionsEntity, err := instancePartitionDao.GetInstancePartitionsByInstanceId(logger, command.InstanceID)
	if instancePartitionsEntity == nil {
		logger.Warnf("FormatPartitionsProcessor GetInstancePartitionsByInstanceId error, sn:%s, instance_id:%s, err:%+v", command.Sn, command.InstanceID, err)
	}

	formatPartitions := agentTypes.FormatPartitions{
		Sn:      command.Sn,
		Version: "2.0",
	}
	ps := []agentTypes.FPartition{}
	for _, v := range instancePartitionsEntity {
		if v.PartitionType == "data" { //数据盘
			continue
		}
		ps = append(ps, agentTypes.FPartition{
			Volume: v.InstancePartitionID,
			FsType: v.PartitionFsType,
			Label:  v.PartitionLabel,
			IsRoot: v.PartitionType == "root",
		})
	}
	formatPartitions.Partitions = ps

	if err := rabbitMq.SendToAgent(command.Sn, formatPartitions); err != nil {
		logger.Warnf("FormatPartitionsProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(formatPartitions), err.Error())
		return
	}
	logger.Infof("FormatPartitionsProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(formatPartitions))

}

func (b FormatPartitionsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
