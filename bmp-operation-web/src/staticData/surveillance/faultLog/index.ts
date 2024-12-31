import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class FalutLogStaticData {

    static faultLogTipData: string[] = [
        'logidTip',
        'snTip',
        'idcNameTip',
        'alertLevelTip',
        'alertTypeTip',
        'alertPartTip',
        'iloIpTip',
        'alertDescTip',
        'alertContentTip',
        'statusTip',
        'countTip',
        'collectTimeTip',
        'eventTimeTip',
        'updateTimeTip'
    ];

    static faultDisposeOfConfirmTip: string[] = [
        'snTip',
        'alertNameTip',
        'alertTypeTip',
        'alertDescTip',
        'collectTimeTip'
    ];

    static faultLogInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            logid: status2,
            sn: status2,
            idcName: status1,
            alertLevel: status1,
            alertType: status1,
            alertPart: status1,
            iloIp: status1,
            alertDesc: status1,
            alertContent: status1,
            status: status1,
            count: status1,
            collectTime: status1,
            eventTime: status1,
            updateTime: status1,
            operate: status2
        }
    };

    static faultLogList = [
        {
            name: global.t('faultLogList.table.label.logId'),
            selected: true,
            disabled: true,
            label: 'logid',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.sn'),
            selected: true,
            disabled: true,
            label: 'sn',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.idc'),
            selected: true,
            disabled: false,
            label: 'idcName',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.level'),
            selected: true,
            disabled: false,
            label: 'alertLevel',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.type'),
            selected: true,
            disabled: false,
            label: 'alertType',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.accessory'),
            selected: true,
            disabled: false,
            label: 'alertPart',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.iloip'),
            selected: true,
            disabled: false,
            label: 'iloIp',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.desc'),
            selected: true,
            disabled: false,
            label: 'alertDesc',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.originalLog'),
            selected: true,
            disabled: false,
            label: 'alertContent',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.status'),
            selected: true,
            disabled: false,
            label: 'status',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.numberOfAlarms'),
            selected: true,
            disabled: false,
            label: 'count',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.faultDetectionTime'),
            selected: true,
            disabled: false,
            label: 'collectTime',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.faultAlarmTime'),
            selected: true,
            disabled: false,
            label: 'eventTime',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.updateTime'),
            selected: true,
            disabled: false,
            label: 'updateTime',
            showTooltip: false
        },
        {
            name: global.t('faultLogList.table.label.operate'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];
};

export default FalutLogStaticData;
