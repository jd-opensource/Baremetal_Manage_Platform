/**
 * @file
 * @author
*/
import {RuleFormType} from '../type';

class RuleFormOpt {
    ruleForm: RuleFormType = reactive<RuleFormType>({
        cpu: 'avg',
        mem: 'avg',
        memUsage: 'avg',
        diskUsage: 'avg',
        diskUsageRate: 'avg',
        diskReadWriteTraffic: 'avg',
        diskReadWriteIOPS: 'avg',
        netWorkBps: 'avg',
        netWorkPackagesNum: 'avg',
        averageLoad1Min: 'avg',
        averageLoad5Min: 'avg',
        averageLoad15Min: 'avg',
        totalTCPConnections: 'avg',
        normalTCPConnections: 'avg',
        totalNumberOfProcesses: 'avg'
    })
}

export default RuleFormOpt;
