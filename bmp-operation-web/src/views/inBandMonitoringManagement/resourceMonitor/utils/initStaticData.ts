/**
 * @file
 * @author
*/
import {language} from 'utils/publicClass.ts';
import {CurrencyType8} from '@utils/publicType';
import {EchartsDataType, EchartsParamsDataType} from '../type';

type NewEchartsDataStaticType = Omit<EchartsDataType, 'filter' | 'some'>[]

class InitStaticData {
    static resourceData: string[] = [language.t('resourceMonitor.search.count.resourceType')];
    static dimensionsData: string[] = [language.t('resourceMonitor.search.count.dimensions')];
    static echartsData: NewEchartsDataStaticType = [
        {
            typeVal: 'cpu',
            id: 'main',
            data: [],
            type: language.t('monitorEcharts.echartsCount.cpu'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.cpu.util',
            unit: '%',
            hoverUnit: '%'
        },
        {
            typeVal: 'mem',
            id: 'main2',
            data: [],
            type: language.t('monitorEcharts.echartsCount.mem'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.memory.util',
            unit: '%',
            hoverUnit: '%'
        },
        {
            typeVal: 'memUsage',
            id: 'main3',
            data: [],
            type: language.t('monitorEcharts.echartsCount.memUsage'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.memory.used',
            unit: 'GB',
            hoverUnit: 'GB'
        },
        {
            typeVal: 'diskUsage',
            id: 'main4',
            data: [],
            type: language.t('monitorEcharts.echartsCount.diskUsage'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.disk.used',
            // metricName: 'bmp.disk.partition.used',
            unit: 'GB',
            hoverUnit: 'GB',
            status: ['diskVal', 'mountpointVal']
        },
        {
            typeVal: 'diskUsageRate',
            id: 'main5',
            data: [],
            type: language.t('monitorEcharts.echartsCount.diskUsageRate'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.disk.util',
            unit: '%',
            hoverUnit: '%',
            status: ['diskVal', 'mountpointVal']
        },
        {
            typeVal: 'diskReadWriteTraffic',
            id: 'main6',
            data: [],
            type: [
                language.t('monitorEcharts.echartsCount.diskReadTraffic'),
                language.t('monitorEcharts.echartsCount.diskWriteTraffic')
            ],
            timeData: [],
            aggregationMethod: 'avg',
            metricName: ['bmp.disk.bytes.read', 'bmp.disk.bytes.write'],
            // unit: 'M',
            unit: 'custom',
            // hoverUnit: 'Bps',
            hoverUnit: '',
            status: ['diskVal']
        },
        {
            typeVal: 'diskReadWriteIOPS',
            id: 'main7',
            data: [],
            type: [
                language.t('monitorEcharts.echartsCount.diskReadIOPS'),
                language.t('monitorEcharts.echartsCount.diskWriteIOPS')
            ],
            timeData: [],
            aggregationMethod: 'avg',
            metricName: ['bmp.disk.counts.read', 'bmp.disk.counts.write'],
            unit: '次/秒',
            hoverUnit: '次/秒',
            status: ['diskVal']
        },
        {
            typeVal: 'netWorkBps',
            id: 'main8',
            data: [],
            type: [
                language.t('monitorEcharts.echartsCount.netWorkIn'),
                language.t('monitorEcharts.echartsCount.netWorkOut')
            ],
            timeData: [],
            aggregationMethod: 'avg',
            metricName: ['bmp.network.bytes.ingress', 'bmp.network.bytes.egress'],
            // unit: 'K',
            // M
            // unit: '',
            unit: 'custom',
            // Kbps
            hoverUnit: '',
            status: ['nicVal']
        },
        {
            typeVal: 'netWorkPackagesNum',
            id: 'main9',
            data: [],
            type: [
                language.t('monitorEcharts.echartsCount.netWorkInPackageNum'),
                language.t('monitorEcharts.echartsCount.netWorkOutPackageNum')
            ],
            timeData: [],
            aggregationMethod: 'avg',
            metricName: ['bmp.network.packets.ingress', 'bmp.network.packets.egress'],
            unit: '个/秒',
            hoverUnit: '个/秒',
            status: ['nicVal']
        },
        {
            typeVal: 'averageLoad1Min',
            id: 'main10',
            data: [],
            type: language.t('monitorEcharts.echartsCount.averageLoad1Min'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.avg.load1',
            unit: '',
            hoverUnit: ''
        },
        {
            typeVal: 'averageLoad5Min',
            id: 'main11',
            data: [],
            type: language.t('monitorEcharts.echartsCount.averageLoad5Min'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.avg.load5',
            unit: '',
            hoverUnit: ''
        },
        {
            typeVal: 'averageLoad15Min',
            id: 'main12',
            data: [],
            type: language.t('monitorEcharts.echartsCount.averageLoad15Min'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.avg.load15',
            unit: '',
            hoverUnit: ''
        },
        {
            typeVal: 'totalTCPConnections',
            id: 'main13',
            data: [],
            type: language.t('monitorEcharts.echartsCount.totalTCPConnections'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.tcp.connect.total',
            unit: '个',
            hoverUnit: '个'
        },
        {
            typeVal: 'normalTCPConnections',
            id: 'main14',
            data: [],
            type: language.t('monitorEcharts.echartsCount.normalTCPConnections'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.tcp.connect.established',
            unit: '个',
            hoverUnit: '个'
        },
        {
            typeVal: 'totalNumberOfProcesses',
            id: 'main15',
            data: [],
            type: language.t('monitorEcharts.echartsCount.totalNumberOfProcesses'),
            timeData: [],
            aggregationMethod: 'avg',
            metricName: 'bmp.process.total',
            unit: '个',
            hoverUnit: '个'
        }
    ];

    static radioData: CurrencyType8[] = [
        {
            label: language.t('monitorEcharts.tabs.h1'),
            value: '1min'
        },
        {
            label: language.t('monitorEcharts.tabs.h6'),
            value: '5min'
        },
        {
            label: language.t('monitorEcharts.tabs.h12'),
            value: '10 min'
        },
        {
            label: language.t('monitorEcharts.tabs.d1'),
            value: '20 min'
        },
        {
            label: language.t('monitorEcharts.tabs.d3'),
            value: '60min'
        },
        {
            label: language.t('monitorEcharts.tabs.d7'),
            value: '120min'
        },
        {
            label: language.t('monitorEcharts.tabs.d14'),
            value: '180min'
        }
    ];

    static echartsParamsData: EchartsParamsDataType = [
        {
            title: language.t('monitorEcharts.echartsCount.cpuUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'cpu',
            id: 'main',
            bigClickVal: {
                type: 'cpu',
                title: language.t('monitorEcharts.echartsCount.cpuUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.memUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'mem',
            id: 'main2',
            bigClickVal: {
                type: 'mem',
                title: language.t('monitorEcharts.echartsCount.memUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.memUsageUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'memUsage',
            id: 'main3',
            bigClickVal: {
                type: 'memUsage',
                title: language.t('monitorEcharts.echartsCount.memUsageUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.diskUsageUnit'),
            showType: [
                language.t('monitorEcharts.radio.instance'),
                language.t('monitorEcharts.radio.disk'),
                language.t('monitorEcharts.radio.diskPartition')
            ],
            model: '',
            type: 'diskUsage',
            id: 'main4',
            bigClickVal: {
                type: 'diskUsage',
                title: language.t('monitorEcharts.echartsCount.diskUsageUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.diskUsageRateUnit'),
            showType: [
                language.t('monitorEcharts.radio.instance'),
                language.t('monitorEcharts.radio.disk'),
                language.t('monitorEcharts.radio.diskPartition')
            ],
            model: '',
            type: 'diskUsageRate',
            id: 'main5',
            bigClickVal: {
                type: 'diskUsageRate',
                title: language.t('monitorEcharts.echartsCount.diskUsageRateUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.diskReadWriteTrafficUnit'),
            showType: [language.t('monitorEcharts.radio.instance'), language.t('monitorEcharts.radio.disk')],
            model: '',
            type: 'diskReadWriteTraffic',
            id: 'main6',
            bigClickVal: {
                type: 'diskReadWriteTraffic',
                title: language.t('monitorEcharts.echartsCount.diskReadWriteTrafficUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.diskReadWriteIOPSUnit'),
            showType: [language.t('monitorEcharts.radio.instance'), language.t('monitorEcharts.radio.disk')],
            model: '',
            type: 'diskReadWriteIOPS',
            id: 'main7',
            bigClickVal: {
                type: 'diskReadWriteIOPS',
                title: language.t('monitorEcharts.echartsCount.diskReadWriteIOPSUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.netWorkBpsUnit'),
            showType: [language.t('monitorEcharts.radio.instance'), language.t('monitorEcharts.radio.netWorkInterface')],
            model: '',
            type: 'netWorkBps',
            id: 'main8',
            bigClickVal: {
                type: 'netWorkBps',
                title: language.t('monitorEcharts.echartsCount.netWorkBpsUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.netWorkPackagesNumUnit'),
            showType: [language.t('monitorEcharts.radio.instance'), language.t('monitorEcharts.radio.netWorkInterface')],
            model: '',
            type: 'netWorkPackagesNum',
            id: 'main9',
            bigClickVal: {
                type: 'netWorkPackagesNum',
                title: language.t('monitorEcharts.echartsCount.netWorkPackagesNumUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.averageLoad1MinUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'averageLoad1Min',
            id: 'main10',
            bigClickVal: {
                type: 'averageLoad1Min',
                title: language.t('monitorEcharts.echartsCount.averageLoad1MinUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.averageLoad5MinUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'averageLoad5Min',
            id: 'main11',
            bigClickVal: {
                type: 'averageLoad5Min',
                title: language.t('monitorEcharts.echartsCount.averageLoad5MinUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.averageLoad15MinUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'averageLoad15Min',
            id: 'main12',
            bigClickVal: {
                type: 'averageLoad15Min',
                title: language.t('monitorEcharts.echartsCount.averageLoad15MinUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.totalTCPConnectionsUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'totalTCPConnections',
            id: 'main13',
            bigClickVal: {
                type: 'totalTCPConnections',
                title: language.t('monitorEcharts.echartsCount.totalTCPConnectionsUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.normalTCPConnectionsUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'normalTCPConnections',
            id: 'main14',
            bigClickVal: {
                type: 'normalTCPConnections',
                title: language.t('monitorEcharts.echartsCount.normalTCPConnectionsUnit'),
                model: ''
            }
        },
        {
            title: language.t('monitorEcharts.echartsCount.totalNumberOfProcessesUnit'),
            showType: [language.t('monitorEcharts.radio.instance')],
            model: '',
            type: 'totalNumberOfProcesses',
            id: 'main15',
            bigClickVal: {
                type: 'totalNumberOfProcesses',
                title: language.t('monitorEcharts.echartsCount.totalNumberOfProcessesUnit'),
                model: ''
            }
        }
    ]
}

export default InitStaticData;
