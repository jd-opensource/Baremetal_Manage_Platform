package scheduler

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/auditLogsDao"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

//onst LOG_DIR = "./log"

func doCleanAuditLogCron() error {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-scheduler-cleanAuditLogCron.log")
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	logger.Point("task_type", "cleanAuditLogCron")
	//没有requestid，临时生成
	logger.Point("logid", commonUtil.GenerateRandUuid())
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	logger.Info("doCleanAuditLogCron starting......")

	stime := time.Now().AddDate(0, 0, -90).Unix()
	q := map[string]interface{}{
		"created_time.lte": stime,
		"is_del":           0,
	}
	u := map[string]interface{}{
		"is_del": 1,
	}
	if err := auditLogsDao.UpdateMultiAuditLogs(logger, q, u); err != nil {
		logger.Warnf("doCleanAuditLogCron.UpdateMultiAuditLogs error, stime:%d, error:%s", stime, err.Error())
		return err
	}

	return nil
}
