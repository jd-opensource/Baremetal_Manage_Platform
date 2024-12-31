package scheduler

import (
	"coding.jd.com/aidc-bmp/oob-log-agent/object"
	"coding.jd.com/aidc-bmp/oob-log-agent/parser"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Job concrete type to execute monitor job
type Job struct {
	cps object.CPSRecord
}

// Run get CPS info
func (j Job) Run(logger *log.Logger) {

	logger.Infof(">>>>>>>> start Job.Run(), sn:%s", j.cps.SN)
	defer logger.Infof(j.cps.SN, "<<<<<<<<end Job.Run(), sn:%s", j.cps.SN)

	// 对iLoIP为空的，无法连接取信息，记录原因，返回
	if j.cps.IloIP == "" {
		logger.Warnf("%+v. CPS iLO IP is blank. Skip", j.cps)
		return
	}

	p := parser.GetParser(logger, j.cps)
	if err := p.ParseFRU(logger); err != nil {
		return
	}
	p.ParseLog(logger)

}
