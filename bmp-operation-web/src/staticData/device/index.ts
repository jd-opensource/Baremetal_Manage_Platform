/**
 * @file
 * @author
*/

import i18n from 'lib/i18n/index.ts'; // 国际化
import DeviceType from './type';
import {CurrencyStatusType, ListShowTooltipType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class DeviceStaticData {
    // 管理状态filter数据
    static manageMentStatus: DeviceType[] = [
        {
            filterParams: 'in',
            text: global.t('deviceList.managementStatus.warehousing'),
            value: 1
        },
        {
            filterParams: 'putaway',
            text: global.t('deviceList.managementStatus.onTheShelf'),
            value: 2
        },
        {
            filterParams: 'created',
            text: global.t('deviceList.managementStatus.created'),
            value: 3
        },
        {
            filterParams: 'putawaying',
            text: global.t('deviceList.managementStatus.onShelf'),
            value: 4
        },
        {
            filterParams: 'removed',
            text: global.t('deviceList.managementStatus.remove'),
            value: 5
        },
        {
            filterParams: 'putawayfail',
            text: global.t('deviceList.managementStatus.error'),
            value: 6
        }
    ];

    static operationalFeedbackData = [
        {
            filterParams: 'success',
            text: global.t('operate.success'),
            value: 1
        },
        {
            filterParams: 'fail',
            text: global.t('operate.fail'),
            value: 2
        },
        {
            filterParams: 'doing',
            text: global.t('operate.doing'),
            value: 3
        }
    ];

    static collectStatusData = [
        {
            filterParams: '1',
            text: global.t('deviceList.status.collected'),
            value: 1
        },
        {
            filterParams: '2',
            text: global.t('deviceList.status.noCollected'),
            value: 2
        },
        {
            filterParams: '3',
            text: global.t('deviceList.status.collecting'),
            value: 3
        },
        {
            filterParams: '4',
            text: global.t('deviceList.status.collectionFailed'),
            value: 4
        }
    ];

    static modelInfoTipData: string[] = [
        'newDisksValTip'
    ];

    // 中间态1
    static intermediate1: string[] = ['creating', 'starting', 'stopping', 'Rebooting' , 'restarting', 'resetting_password', 'destroying', 'upgrading', 'reinstalling', 'resettingPassword'];

    // 中间态2
    static intermediate2: string[] = ['Creating', 'Starting up', 'Shutting Down', 'Rebooting', 'Restarting', 'Reset Password', 'Destroying', 'Destructing', 'Destroying', 'Upgrading', 'Reinstalling System'];

    static deviceList = [
        {
            name: global.t('deviceList.content.label.sn'),
            selected: true,
            disabled: true,
            label: 'sn',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.name'),
            selected: true,
            disabled: false,
            label: 'deviceSeries',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.modelName'),
            selected: true,
            disabled: false,
            label: 'deviceTypeName',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.managementStatus'),
            selected: true,
            disabled: false,
            label: 'manageStatusName',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.collectionStatus'),
            selected: true,
            disabled: false,
            label: 'collectStatus',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.computerRoomName'),
            selected: true,
            disabled: false,
            label: 'idcName',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.cabinetCode'),
            selected: true,
            disabled: false,
            label: 'cabinet',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.uBit'),
            selected: true,
            disabled: false,
            label: 'uPosition',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.brand'),
            selected: true,
            disabled: false,
            label: 'brand',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.model'),
            selected: true,
            disabled: false,
            label: 'model',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.architecture'),
            selected: true,
            disabled: false,
            label: 'architecture',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.instanceImage'),
            selected: true,
            disabled: false,
            label: 'imageName',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.instanceName'),
            selected: true,
            disabled: false,
            label: 'instanceName',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.instanceID'),
            selected: true,
            disabled: false,
            label: 'instanceId',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.instanceStatus'),
            selected: true,
            disabled: false,
            label: 'instanceStatus',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.lockStatus'),
            selected: true,
            disabled: false,
            label: 'instanceLockedStatus',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.instanceOwner'),
            selected: true,
            disabled: false,
            label: 'instanceOwer',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.CPU'),
            selected: true,
            disabled: false,
            label: 'cpuInfo',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.storage'),
            selected: true,
            disabled: false,
            label: 'memInfo',
            showTooltip: false
        },
        // {
        //     name: global.t('deviceList.content.label.sysDisc'),
        //     selected: true,
        //     disabled: false,
        //     label: 'svInfo',
        //     showTooltip: false
        // },
        // {
        //     name: global.t('deviceList.content.label.dataDisc'),
        //     selected: true,
        //     disabled: false,
        //     label: 'dvInfo',
        //     showTooltip: false
        // },
        {
            name: global.t('deviceList.content.label.GPU'),
            selected: true,
            disabled: false,
            label: 'gpuInfo',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.switchIPNetworkPortOne'),
            selected: true,
            disabled: false,
            label: 'switchIp1',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.switchIPNetworkPortTwo'),
            selected: true,
            disabled: false,
            label: 'switchIp2',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.networkPortOneUplinkPort'),
            selected: true,
            disabled: false,
            label: 'switchPort1',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.networkPortTwoUplinkPort'),
            selected: true,
            disabled: false,
            label: 'switchPort2',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.networkSettings'),
            selected: true,
            disabled: false,
            label: 'nicInfo',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.outOfBandIP'),
            selected: true,
            disabled: false,
            label: 'iloIp',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.intranceIPv4'),
            selected: true,
            disabled: false,
            label: 'privateIpv4',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.intranceIPv4First'),
            selected: true,
            disabled: false,
            label: 'privateEth1Ipv4',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.intranetIPv6'),
            selected: true,
            disabled: false,
            label: 'privateIpv6',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.intranetIPv6First'),
            selected: true,
            disabled: false,
            label: 'privateEth1Ipv6',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.createTime'),
            selected: true,
            disabled: false,
            label: 'createdTime',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.remark'),
            selected: true,
            disabled: false,
            label: 'description',
            showTooltip: false
        },
        {
            name: global.t('deviceList.content.label.operate.name'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];

    static auditLogsList = [
        {
            name: global.t('deviceDetail.operatLog.table.label.logId'),
            selected: true,
            disabled: true,
            label: 'logid',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.operatLog.table.label.operateName'),
            selected: true,
            disabled: true,
            label: 'operationName',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.operatLog.table.label.operatePeople'),
            selected: true,
            disabled: false,
            label: 'userName',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.operatLog.table.label.operateTime'),
            selected: true,
            disabled: false,
            label: 'operateTime',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.operatLog.table.label.operateFeedback'),
            selected: true,
            disabled: false,
            label: 'result',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.operatLog.table.label.failCode'),
            selected: true,
            disabled: false,
            label: 'failReason',
            showTooltip: false
        }
    ];

    static deviceFaultLogList = [
        {
            name: global.t('deviceDetail.table.label.serialNumber'),
            selected: true,
            disabled: true,
            label: 'logid',
            showTooltip: false
        },
        {
            name:  global.t('deviceDetail.table.label.faultType'),
            selected: true,
            disabled: false,
            label: 'alertType',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.table.label.faultDesc'),
            selected: true,
            disabled: false,
            label: 'alertDesc',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.table.label.faultAlarmTime'),
            selected: true,
            disabled: false,
            label: 'eventTime',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.table.label.updateTime'),
            selected: true,
            disabled: false,
            label: 'updateTime',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.table.label.alarmNum'),
            selected: true,
            disabled: false,
            label: 'count',
            showTooltip: false
        },
        {
            name: global.t('deviceDetail.table.label.status'),
            selected: true,
            disabled: true,
            label: 'status',
            showTooltip: false
        },
    ];

    static deviceFaultLogInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            logid: status2,
            alertType: status1,
            alertDesc: status1,
            eventTime: status1,
            updateTime: status1,
            count: status1,
            status: status2,
        }
    };

    static deviceInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            sn: status2,
            deviceTypeName: status1,
            deviceSeries: status1,
            manageStatusName: status1,
            collectStatus: status1,
            idcName: status1,
            cabinet: status1,
            uPosition: status1,
            brand: status1,
            model: status1,
            architecture: status1,
            imageName: status1,
            instanceName: status1,
            instanceId: status1,
            instanceStatus: status1,
            instanceLockedStatus: status1,
            instanceOwer: status1,
            cpuInfo: status1,
            memInfo: status1,
            // svInfo: status1,
            // dvInfo: status1,
            gpuInfo: status1,
            switchIp1: status1,
            switchIp2: status1,
            switchPort1: status1,
            switchPort2: status1,
            nicInfo: status1,
            iloIp: status1,
            privateIpv4: status1,
            privateEth1Ipv4: status1,
            privateIpv6: status1,
            privateEth1Ipv6: status1,
            createdTime: status1,
            operate: status2
        }
    };

    static auditLogsInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) => {
        return {
            logid: status2,
            operationName: status2,
            userName: status1,
            operateTime: status1,
            result: status1,
            failReason: status1
        }
    };

    static deviceTipData: string[] = [
        'snTip',
        'deviceTypeNameTip',
        'idcNameTip',
        'cabinetTip',
        'uPositionTip',
        'brandTip',
        'modelTip',
        'architectureTip',
        'imageNameTip',
        'instanceNameTip',
        'instanceIdTip',
        'userNameTip',
        'cpuInfoTip',
        'memInfoTip',
        'svInfoTip',
        'dvInfoTip',
        'gpuInfoTip',
        'switchIp1Tip',
        'switchIp2Tip',
        'switchPort1Tip',
        'switchPort2Tip',
        'interfaceModeNameTip',
        'iloIpTip',
        'privateIpv4Tip',
        'privateEth1Ipv4Tip',
        'privateIpv6Tip',
        'privateEth1Ipv6Tip',
        'descriptionTip'
    ];

    static batchEditInstanceNameTipData: string[] = [
        'instanceNameTip',
        'instanceIdTip'
    ];

    static resetSysTemKeyTooltip: ListShowTooltipType = {
        name: {
            showTooltip: false
        }
    };

    static deviceFaultLogTipData: string[] = [
        'alertTypeTip',
        'alertDescTip'
    ];

    static loadingStatusData: string[] = ['creating', 'Rebooting', 'restarting', 'starting', 'stopping', 'destroying', 'reinstalling', 'resetting_password'];

    static FailedData: string[] = [
        'Creation failed',
        'Startup failed',
        'Shutdown failed',
        'Reboot failed',
        'Delete failed',
        'Reinstall failed',
        'Resetpasswd failed'
    ];

    static recoveryInstance: string[] = ['Creation failed', 'stopped'];
    
    static lockUnlocked: string[] = ['Creating', 'creating', 'Creation failed', 'destroying' , 'Destroying', 'Delete failed', 'Reinstall failed', 'Resetpasswd failed'];

    static resetInstanceStatus: string[] = ['Startup failed', 'Shutdown failed', 'Reboot failed', 'Delete failed'];

    static deleteDevice: string[] = ['in', 'putawayfail', 'removed'];

    static resetInstance: string[] = [
        'Startup failed',
        'Shutdown failed',
        'Reboot failed',
        'Delete failed'
    ];
};

export default DeviceStaticData;
