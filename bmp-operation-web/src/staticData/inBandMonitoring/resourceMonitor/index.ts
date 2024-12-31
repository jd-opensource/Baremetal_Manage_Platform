/**
 * @file
 * @author
*/

import i18n from 'lib/i18n/index.ts'; // 国际化

// 国际化
const {global} = i18n;

class ResourceMonitorStaticData {
    static monitoringIndicators = [
        {
            name: global.t('monitorEcharts.echartsCount.cpu'),
            value: 'bmp.cpu.util'
        },
        {
            name: global.t('monitorEcharts.echartsCount.mem'),
            value: 'bmp.memory.util'
        },
        {
            name: global.t('monitorEcharts.echartsCount.memUsage'),
            value: 'bmp.memory.used'
        },
        {
            name: global.t('monitorEcharts.echartsCount.diskUsage'),
            value: 'bmp.disk.used'
            // value: 'bmp.disk.partition.used'
        },
        {
            name: global.t('monitorEcharts.echartsCount.diskUsageRate'),
            value: 'bmp.disk.util'
        },
        {
            name: global.t('monitorEcharts.echartsCount.diskReadWriteTraffic'),
            value: 'bmp.disk.bytes.read,bmp.disk.bytes.write'
        },
        {
            name: global.t('monitorEcharts.echartsCount.diskReadWriteIOPS'),
            value: 'bmp.disk.counts.read,bmp.disk.counts.write '
        },
        {
            name: global.t('monitorEcharts.echartsCount.netWorkBps'),
            value: 'bmp.network.bytes.ingress,bmp.network.bytes.egress'
        },
        {
            name: global.t('monitorEcharts.echartsCount.netWorkPackagesNumber'),
            value: 'bmp.network.packets.ingress,bmp.network.packets.egress'
        },
        {
            name: global.t('monitorEcharts.echartsCount.averageLoad1Min'),
            value: 'bmp.avg.load1'
        },
        {
            name: global.t('monitorEcharts.echartsCount.averageLoad5Min'),
            value: 'bmp.avg.load5'
        },
        {
            name: global.t('monitorEcharts.echartsCount.averageLoad15Min'),
            value: 'bmp.avg.load15'
        },
        {
            name: global.t('monitorEcharts.echartsCount.totalTCPConnections'),
            value: 'bmp.tcp.connect.total'
        },
        {
            name: global.t('monitorEcharts.echartsCount.normalTCPConnections'),
            value: 'bmp.tcp.connect.established'
        },
        {
            name: global.t('monitorEcharts.echartsCount.totalNumberOfProcesses'),
            value: 'bmp.process.total'
        }
    ]
}

export default ResourceMonitorStaticData;
